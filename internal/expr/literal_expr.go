package expr

import (
	"reflect"
)

// LiteralExpr represents a literal expression that evaluates to a
// specific value.
type LiteralExpr reflect.Value

// Literal creates a new LiteralExpr with the given value.
func Literal(v any) LiteralExpr {
	return LiteralExpr(reflect.ValueOf(v))
}

// Eval evaluates the literal expression and returns its value.
func (e LiteralExpr) Eval(*Context) (reflect.Value, error) {
	rv := reflect.Value(e)
	if !rv.IsValid() {
		return reflect.Value{}, nil
	}
	return rv, nil
}

var _ Expr = (*LiteralExpr)(nil)
