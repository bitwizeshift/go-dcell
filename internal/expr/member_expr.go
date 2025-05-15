package expr

import (
	"reflect"
	"slices"

	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// MemberExpr is an expression that accesses a member field of a struct or a
// map key. If the input context is a nil value, the output will also be a
// nil value.
//
// If the field does not exist in the input context, an [errs.NameError] is
// returned.
type MemberExpr string

// Member returns a [MemberExpr] with the given name.
func Member(name string) MemberExpr {
	return MemberExpr(name)
}

// Eval evaluates the member expression. It returns the value of the member
// field in the current context.
func (e MemberExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv := reflectconv.Deref(ctx.Current)
	if reflectconv.IsNil(rv) {
		return reflect.Value{}, nil
	}

	rv = reflectconv.Deref(rv)
	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Map:
		return e.evalMap(rv)
	case reflect.Struct:
		return e.evalStruct(rv, rt)
	case reflect.Slice, reflect.Array:
		return e.evalSlice(rv)
	}
	return reflect.Value{}, errs.NewNameError(string(e), slices.Values([]string{}))
}

func (e MemberExpr) evalMap(rv reflect.Value) (reflect.Value, error) {
	if rv.IsNil() {
		return reflect.Value{}, nil
	}
	keyType := rv.Type().Key()
	if keyType.Kind() != reflect.String {
		return reflect.Value{}, errs.NewNameError(string(e), slices.Values([]string{}))
	}

	key := reflect.ValueOf(string(e)).Convert(keyType)
	value := rv.MapIndex(key)
	if !value.IsValid() {
		keys := e.mapKeys(rv)
		return reflect.Value{}, errs.NewNameError(string(e), slices.Values(keys))
	}

	return value, nil
}

// mapKeys returns a slice of string keys from the map value. This expects that
// the map's key has an underlying kind of string.
func (e MemberExpr) mapKeys(rv reflect.Value) []string {
	keys := rv.MapKeys()
	var result []string
	for _, key := range keys {
		result = append(result, key.String())
	}
	return result
}

func (e MemberExpr) evalStruct(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	numFields := rt.NumField()
	if numFields == 0 {
		return reflect.Value{}, nil
	}

	var names []string
	for i := range numFields {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}

		tag := field.Tag.Get("dcell")
		if tag == "" {
			tag = field.Name
		}
		if tag == string(e) {
			rfield := rv.Field(i)
			return rfield, nil
		}
		names = append(names, tag)
	}
	return reflect.Value{}, errs.NewNameError(string(e), slices.Values(names))
}

func (e MemberExpr) evalSlice(rv reflect.Value) (reflect.Value, error) {
	var entries []reflect.Value
	for i := range rv.Len() {
		entry := rv.Index(i)
		entry = reflectconv.Deref(entry)
		switch entry.Kind() {
		case reflect.Map:
			value, err := e.evalMap(entry)
			if err != nil {
				return reflect.Value{}, err
			}
			entries = append(entries, value)
		case reflect.Struct:
			value, err := e.evalStruct(entry, entry.Type())
			if err != nil {
				return reflect.Value{}, err
			}
			entries = append(entries, value)
		default:
			return reflect.Value{}, errs.NewNameError(string(e), slices.Values([]string{}))
		}
	}
	if len(entries) == 0 {
		return reflect.Value{}, nil
	}
	sliceType := e.computeSliceType(entries)
	slice := reflect.MakeSlice(sliceType, 0, len(entries))
	for _, entry := range entries {
		slice = reflect.Append(slice, entry)
	}
	return slice, nil
}

func (e MemberExpr) computeSliceType(entries []reflect.Value) reflect.Type {
	current := entries[0].Type()
	for _, field := range entries[1:] {
		if field.Type() != current {
			return reflect.TypeFor[[]any]()
		}
	}
	return reflect.SliceOf(current)
}

var _ Expr = (*MemberExpr)(nil)
