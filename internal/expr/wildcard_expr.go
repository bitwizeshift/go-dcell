package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// WildcardExpr is a wildcard expression that matches any field of a struct
// that contains a `dcell` struct tag, or any value associated to a key in a
// map.
//
// If the current context input is a nil value, the result is nil.
// If the current context input is not a struct or a map, an
// error is returned.
type WildcardExpr struct{}

// Wildcard returns a wildcard expression.
func Wildcard() WildcardExpr {
	return WildcardExpr{}
}

// Eval evaluates the wildcard expression.
func (e WildcardExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv := ctx.Current
	if reflectconv.IsNil(rv) {
		return reflect.Value{}, nil
	}

	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return reflect.Value{}, nil
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Struct:
		return e.evalStruct(rv), nil
	case reflect.Map:
		return e.evalMap(rv), nil
	}

	return reflect.Value{}, fmt.Errorf("wildcard: '*' only usable on struct and map, got %s", rv.Type().Name())
}

func (e WildcardExpr) evalStruct(rv reflect.Value) reflect.Value {
	fields := e.extractFields(rv)
	if len(fields) == 0 {
		return reflect.Value{}
	}
	slice := reflect.MakeSlice(e.fieldSliceType(fields), 0, len(fields))
	for _, field := range fields {
		slice = reflect.Append(slice, field)
	}

	return slice
}

func (e WildcardExpr) extractFields(rv reflect.Value) []reflect.Value {
	var result []reflect.Value
	rt := rv.Type()
	for i := range rt.NumField() {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		result = append(result, rv.Field(i))
	}
	return result
}

func (e WildcardExpr) fieldSliceType(fields []reflect.Value) reflect.Type {
	current := fields[0].Type()
	for _, field := range fields[1:] {
		if field.Type() != current {
			return reflect.TypeFor[[]any]()
		}
	}
	return reflect.SliceOf(current)
}

func (e WildcardExpr) evalMap(rv reflect.Value) reflect.Value {
	if rv.IsNil() || rv.Len() == 0 {
		return reflect.Value{}
	}
	rt := rv.Type()
	valueType := rt.Elem()
	slice := reflect.MakeSlice(reflect.SliceOf(valueType), 0, rv.Len())
	for _, key := range rv.MapKeys() {
		value := rv.MapIndex(key)
		slice = reflect.Append(slice, value)
	}
	return slice
}

var _ Expr = (*WildcardExpr)(nil)
