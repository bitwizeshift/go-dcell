package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// TernaryExpr represents a ternary expression (condition ? trueExpr : falseExpr).
// This expression requires that Condition evaluates to a boolean value,
// and will only evaluate either TrueExpr or FalseExpr based on whether the
// Condition evaluates to true or false, respectively.
type TernaryExpr struct {
	Condition Expr
	TrueExpr  Expr
	FalseExpr Expr
}

// Ternary creates a new TernaryExpr with the given condition, true expression,
func Ternary(condition, trueExpr, falseExpr Expr) *TernaryExpr {
	return &TernaryExpr{
		Condition: condition,
		TrueExpr:  trueExpr,
		FalseExpr: falseExpr,
	}
}

// Elvis creates a new TernaryExpr that evaluates to the condition itself if
// true, or to the falseExpr if the condition is false.
// This is a short-hand for (condition ? condition : falseExpr).
func Elvis(condition, falseExpr Expr) *TernaryExpr {
	return &TernaryExpr{
		Condition: condition,
		TrueExpr:  condition,
		FalseExpr: falseExpr,
	}
}

// Eval evaluates the ternary expression.
func (e *TernaryExpr) Eval(ctx *Context) (reflect.Value, error) {
	result, err := e.Condition.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsTruthy(result) {
		return e.TrueExpr.Eval(ctx)
	}
	return e.FalseExpr.Eval(ctx)
}

var _ Expr = (*TernaryExpr)(nil)
