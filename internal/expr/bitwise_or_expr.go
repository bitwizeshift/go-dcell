package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type BitwiseOrExpr struct {
	Left, Right Expr
}

func BitwiseOr(left, right Expr) *BitwiseOrExpr {
	return &BitwiseOrExpr{
		Left:  left,
		Right: right,
	}
}

func (e *BitwiseOrExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	is, err := reflectconv.Uint64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := is[0], is[1]
	return reflect.ValueOf(lhsInt | rhsInt), nil
}

var _ Expr = (*BitwiseOrExpr)(nil)

type BitwiseXorExpr struct {
	Left, Right Expr
}

func BitwiseXor(left, right Expr) *BitwiseXorExpr {
	return &BitwiseXorExpr{
		Left:  left,
		Right: right,
	}
}

func (e *BitwiseXorExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	is, err := reflectconv.Uint64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := is[0], is[1]
	return reflect.ValueOf(lhsInt ^ rhsInt), nil
}

var _ Expr = (*BitwiseXorExpr)(nil)
