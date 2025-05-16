package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type LogicalAndExpr struct {
	Left, Right Expr
}

func LogicalAnd(left, right Expr) *LogicalAndExpr {
	return &LogicalAndExpr{
		Left:  left,
		Right: right,
	}
}

func (e *LogicalAndExpr) Eval(ctx *Context) (reflect.Value, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	if !reflectconv.IsTruthy(left) {
		return reflect.ValueOf(false), nil
	}

	right, err := e.Right.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(reflectconv.IsTruthy(right)), nil
}

var _ Expr = (*LogicalAndExpr)(nil)
