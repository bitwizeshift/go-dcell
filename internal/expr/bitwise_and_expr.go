package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// BitwiseAndExpr represents a bitwise AND expression.
// It takes two expressions and evaluates them to perform a bitwise AND
// operation.
type BitwiseAndExpr struct {
	Left, Right Expr
}

// BitwiseAnd returns a [BitwiseAndExpr] with the given left and right operands.
func BitwiseAnd(left, right Expr) *BitwiseAndExpr {
	return &BitwiseAndExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the BitwiseAndExpr.
func (e *BitwiseAndExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}

	is, err := reflectconv.Uint64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := is[0], is[1]
	return reflect.ValueOf(lhsInt & rhsInt), nil
}

var _ Expr = (*BitwiseAndExpr)(nil)
