package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// BitwiseNotExpr represents a bitwise NOT expression.
// It takes a single expression and evaluates it to perform a bitwise NOT (~)
// operation on the result.
type BitwiseNotExpr struct {
	Expr Expr
}

// BitwiseNot returns a [BitwiseNotExpr] with the given operand.
func BitwiseNot(expr Expr) BitwiseNotExpr {
	return BitwiseNotExpr{
		Expr: expr,
	}
}

// Eval evaluates the BitwiseNotExpr.
func (e BitwiseNotExpr) Eval(ctx *Context) (reflect.Value, error) {
	val, err := e.Expr.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	i, err := reflectconv.Uint64(val)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(^i), nil
}

var _ Expr = (*BitwiseNotExpr)(nil)
