package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// evalMultiple evaluates multiple expressions and returns their results as a
// slice of reflect.Value.
func evalMultiple(ctx *Context, exprs ...Expr) ([]reflect.Value, error) {
	var result []reflect.Value
	for _, expr := range exprs {
		val, err := expr.Eval(ctx)
		if err != nil {
			return nil, err
		}
		val = reflectconv.Deref(val)
		result = append(result, val)
	}
	return result, nil
}

func evalTwo(ctx *Context, left, right Expr) (reflect.Value, reflect.Value, error) {
	results, err := evalMultiple(ctx, left, right)
	if err != nil {
		return reflect.Value{}, reflect.Value{}, err
	}
	return results[0], results[1], nil
}
