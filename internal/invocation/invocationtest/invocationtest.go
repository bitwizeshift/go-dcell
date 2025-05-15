package invocationtest

import "reflect"

// AlwaysError returns a function that always returns the given error.
func AlwaysError(err error) func(...reflect.Value) (reflect.Value, error) {
	return func(...reflect.Value) (reflect.Value, error) {
		return reflect.Value{}, err
	}
}

// AlwaysReturn returns a function that always returns the given value.
func AlwaysReturn(v any) func(...reflect.Value) (reflect.Value, error) {
	return func(...reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(v), nil
	}
}
