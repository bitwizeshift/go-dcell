package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// ImpliesExpr represents an implication expression in dcell.
// It is used to evaluate the logical implication of two expressions.
type ImpliesExpr struct {
	Left  Expr
	Right Expr
}

// Implies creates a new ImpliesExpr with the given left and right expressions.
func Implies(left, right Expr) *ImpliesExpr {
	return &ImpliesExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the implication expression.
func (e *ImpliesExpr) Eval(ctx *Context) (reflect.Value, error) {
	left, right, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(reflectconv.IsTruthy(left) == reflectconv.IsTruthy(right)), nil
}
