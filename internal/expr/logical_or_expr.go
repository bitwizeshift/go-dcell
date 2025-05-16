package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type LogicalOrExpr struct {
	Left, Right Expr
}

func LogicalOr(left, right Expr) *LogicalOrExpr {
	return &LogicalOrExpr{
		Left:  left,
		Right: right,
	}
}

func (e *LogicalOrExpr) Eval(ctx *Context) (reflect.Value, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsTruthy(left) {
		return reflect.ValueOf(true), nil
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(reflectconv.IsTruthy(right)), nil
}

var _ Expr = (*LogicalOrExpr)(nil)
