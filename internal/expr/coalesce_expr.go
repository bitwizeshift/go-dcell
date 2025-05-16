package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// CoalesceExpr implements the coalesce operator ?? in expressions.
// This returns the first non-null value from the left or right operand.
// This operation short-circuits, meaning if the left operand is not null,
// the right operand is not evaluated.
type CoalesceExpr struct {
	Left, Right Expr
}

// Coalesce creates a new CoalesceExpr with the given left and right operands.
func Coalesce(left, right Expr) *CoalesceExpr {
	return &CoalesceExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the CoalesceExpr.
func (e *CoalesceExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, err := e.Left.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}

	if !reflectconv.IsNil(lhs) {
		return lhs, nil
	}
	return e.Right.Eval(ctx)
}

var _ Expr = (*CoalesceExpr)(nil)
