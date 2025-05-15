package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// LogicalNotExpr represents a logical NOT expression which negates a
// boolean value.
type LogicalNotExpr struct {
	Expr Expr
}

// LogicalNot creates a new [LogicalNotExpr] that negates the specified expression.
func LogicalNot(expr Expr) LogicalNotExpr {
	return LogicalNotExpr{
		Expr: expr,
	}
}

// Eval evaluates the logical NOT expression.
func (e LogicalNotExpr) Eval(ctx *Context) (reflect.Value, error) {
	got, err := e.Expr.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(!reflectconv.IsTruthy(got)), nil
}

var _ Expr = (*LogicalNotExpr)(nil)
