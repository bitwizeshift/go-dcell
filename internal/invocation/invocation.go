/*
Package invocation provides a mechanism for defining and invoking functions in
a yamlpath expression.
*/
package invocation

import (
	"errors"
	"fmt"
	"iter"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/invocation/arity"
)

// ErrUnknownFunc is an error that is returned when a function is not found in
// the function table.
var (
	ErrUnknownFunc = errors.New("unknown function")
	ErrBadFunc     = errors.New("bad function")
	ErrBadArgument = errors.New("bad argument")
)

// Entry is a function entry in the function table.
type Entry struct {
	fn    funcEntry
	arity arity.Arity
}

// SetArity sets the arity of the function entry.
func (e *Entry) SetArity(a arity.Arity) {
	e.arity = a
}

// TestArity tests the arity of the function with the given number of arguments.
func (e *Entry) TestArity(n int) error {
	return e.arity.Check(n)
}

// Invoke invokes the function with the given context and arguments.
func (e *Entry) Invoke(params ...reflect.Value) (reflect.Value, error) {
	if err := e.TestArity(len(params)); err != nil {
		return reflect.Value{}, err
	}

	return e.fn(params...)
}

// var _ Func = (*Entry)(nil)

// Table is an invocation table that maps function names to function definitions.
type Table struct {
	entries map[string]*Entry

	// parent is the parent table that this table is derived from.
	parent *Table
}

// NewTable creates a new function table.
func NewTable() *Table {
	return &Table{
		entries: make(map[string]*Entry),
	}
}

// New creates a new function table from the current parent table.
func (t *Table) New() *Table {
	return &Table{
		entries: make(map[string]*Entry),
		parent:  t,
	}
}

// Add adds a function to the table.
// By default, functions have an arity of zero -- meaning no arguments may be
// provided to them. To set the arity of the function, use the [Entry.SetArity]
// method on the returned [Entry].
func (t *Table) Add(name string, fn funcEntry) *Entry {
	entry := &Entry{
		fn:    fn,
		arity: arity.None(),
	}
	t.entries[name] = entry
	return entry
}

// Lookup performs a function lookup in the table, and if the function is not
// found, it will recursively search the parent table. If the function is not
// found in the parent table, ok will be set to false and fn will be nil.
func (t *Table) Lookup(name string) (fn *Entry, ok bool) {
	if t == nil || t.entries == nil {
		return nil, false
	}
	entry, ok := t.entries[name]
	if !ok {
		return t.parent.Lookup(name)
	}

	return entry, true
}

// FunctionNames returns an iterator over the function names in the table.
// Names are not guaranteed to be unique if a parent table is used and the
// derived table shadows a function in the parent table.
func (t *Table) FunctionNames() iter.Seq[string] {
	return t.iterate
}

func (t *Table) iterate(yield func(string) bool) {
	for name := range t.entries {
		if !yield(name) {
			return
		}
	}
	if t.parent != nil {
		t.parent.iterate(yield)
	}
}

type funcEntry = func(params ...reflect.Value) (reflect.Value, error)

var errType = reflect.TypeFor[error]()

// AddFunc is a convenience method for adding a normal Go function to the
// function table. This will implicitly convert the function to a
// [funcEntry] and set the arity of the function based on the number of
// arguments and return values of the function. The function must have at
// least one return value, and at most two return values with the second return
// value being an [error] type.
func (t *Table) AddFunc(name string, fn any) error {
	entry, arity, err := t.makeFunc(fn)
	if err != nil {
		return err
	}
	t.Add(name, entry).SetArity(arity)
	return nil
}

func (t *Table) makeFunc(fn any) (funcEntry, arity.Arity, error) {
	rv := reflect.ValueOf(fn)
	rt := rv.Type()
	if err := t.validateFuncType(rv, rt); err != nil {
		return nil, nil, err
	}

	ar := t.getFuncArity(rt)
	collectArgs := t.getCollectFunc(rt)
	getOut := t.getOutputFunc(rt)

	result := func(in ...reflect.Value) (reflect.Value, error) {
		args, err := collectArgs(in)
		if err != nil {
			return reflect.Value{}, err
		}

		out := rv.Call(args)
		return getOut(out)
	}
	return result, ar, nil
}

func (t *Table) validateFuncType(rv reflect.Value, rt reflect.Type) error {
	if rv.Kind() != reflect.Func {
		return fmt.Errorf("%w: expected a function, got %s", ErrBadFunc, rv.Kind())
	}
	if rt.NumOut() > 2 || rt.NumOut() == 0 {
		return fmt.Errorf("%w: expected a function with 1 or 2 return values, got %d", ErrBadFunc, rt.NumOut())
	}
	if rt.NumOut() == 2 && !rt.Out(1).Implements(errType) {
		return fmt.Errorf("%w: second return value must be an error, got %s", ErrBadFunc, rt.Out(1).Name())
	}
	return nil
}

func (t *Table) getFuncArity(rt reflect.Type) arity.Arity {
	if rt.IsVariadic() {
		return arity.AtLeast(rt.NumIn() - 1)
	}
	return arity.Exactly(rt.NumIn())
}

func (t *Table) getCollectFunc(rt reflect.Type) func([]reflect.Value) ([]reflect.Value, error) {
	if rt.IsVariadic() {
		return func(in []reflect.Value) ([]reflect.Value, error) {
			var args []reflect.Value
			for i := range rt.NumIn() - 1 {
				if !in[i].Type().AssignableTo(rt.In(i)) {
					return nil, conversionError(i, rt.In(i), in[i].Type())
				}
				args = append(args, in[i])
			}
			rest := in[rt.NumIn()-1:]
			variadicType := rt.In(rt.NumIn() - 1).Elem()
			for i := range rest {
				if !rest[i].Type().AssignableTo(variadicType) {
					return nil, conversionError(i+rt.NumIn()-1, variadicType, rest[i].Type())
				}
			}
			args = append(args, rest...)
			return args, nil
		}
	}
	return func(in []reflect.Value) ([]reflect.Value, error) {
		for i := range rt.NumIn() {
			if !in[i].Type().AssignableTo(rt.In(i)) {
				return nil, conversionError(i, rt.In(i), in[i].Type())
			}
		}
		return in, nil
	}
}

func (t *Table) getOutputFunc(rt reflect.Type) func([]reflect.Value) (reflect.Value, error) {
	if rt.NumOut() == 2 {
		return func(out []reflect.Value) (reflect.Value, error) {
			if !out[1].IsNil() {
				return reflect.Value{}, out[1].Interface().(error)
			}
			return out[0], nil
		}
	}
	return func(out []reflect.Value) (reflect.Value, error) {
		return out[0], nil
	}
}

func conversionError(i int, want reflect.Type, got reflect.Type) error {
	return fmt.Errorf("%w: argument %d must be of type %s, got %s", ErrBadArgument, i, want.Name(), got.Name())
}
