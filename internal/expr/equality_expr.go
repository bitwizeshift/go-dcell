package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

// EqualityExpr represents an equality expression.
type EqualityExpr struct {
	Left, Right Expr
}

// Equal is a convenience function for creating an equality expression
// that checks if two expressions are equal.
func Equal(left, right Expr) *EqualityExpr {
	return &EqualityExpr{
		Left:  left,
		Right: right,
	}
}

// NotEqual is a convenience function for creating an equality expression
// that checks if two expressions are not equal.
func NotEqual(left, right Expr) Expr {
	return LogicalNot(Equal(left, right))
}

// Eval evaluates the EqualityExpr.
func (e *EqualityExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(reflectcmp.Equal(lhs, rhs)), nil
}

var _ Expr = (*EqualityExpr)(nil)
