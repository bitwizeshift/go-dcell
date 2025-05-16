package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// BitwiseShiftLeftExpr represents a bitwise left shift expression (<<).
//
// This expression will always produce a 64-bit unsigned integer result.
type BitwiseShiftLeftExpr struct {
	Left, Right Expr
}

// BitwiseShiftLeft creates a new BitwiseShiftLeftExpr with the given left and
// right expressions. The left expression is the value to be shifted, and the
// right expression is the number of bits to shift.
func BitwiseShiftLeft(left, right Expr) *BitwiseShiftLeftExpr {
	return &BitwiseShiftLeftExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the bitwise left shift expression.
func (e *BitwiseShiftLeftExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	i, err := reflectconv.Uint64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := i[0], i[1]

	return reflect.ValueOf(lhsInt << rhsInt), nil
}

var _ Expr = (*BitwiseShiftLeftExpr)(nil)

// BitwiseShiftRightExpr represents a bitwise right shift expression (>>).
// This expression will always produce a 64-bit unsigned integer result.
type BitwiseShiftRightExpr struct {
	Left, Right Expr
}

// BitwiseShiftRight creates a new BitwiseShiftRightExpr with the given left and
// right expressions. The left expression is the value to be shifted, and the
// right expression is the number of bits to shift.
func BitwiseShiftRight(left, right Expr) *BitwiseShiftRightExpr {
	return &BitwiseShiftRightExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the bitwise right shift expression.
func (e *BitwiseShiftRightExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	i, err := reflectconv.Uint64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := i[0], i[1]

	return reflect.ValueOf(lhsInt >> rhsInt), nil
}

var _ Expr = (*BitwiseShiftLeftExpr)(nil)
