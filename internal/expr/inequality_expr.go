package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

type InequalityExpr struct {
	Left, Right Expr
	Compare     func(left, right reflect.Value) bool
}

func LessThan(left, right Expr) *InequalityExpr {
	return &InequalityExpr{
		Left:  left,
		Right: right,
		Compare: func(lhs, rhs reflect.Value) bool {
			return reflectcmp.Compare(lhs, rhs) < 0
		},
	}
}

func LessThanOrEqual(left, right Expr) *InequalityExpr {
	return &InequalityExpr{
		Left:  left,
		Right: right,
		Compare: func(lhs, rhs reflect.Value) bool {
			return reflectcmp.Compare(lhs, rhs) <= 0
		},
	}
}
func GreaterThan(left, right Expr) *InequalityExpr {
	return &InequalityExpr{
		Left:  left,
		Right: right,
		Compare: func(lhs, rhs reflect.Value) bool {
			return reflectcmp.Compare(lhs, rhs) > 0
		},
	}
}
func GreaterThanOrEqual(left, right Expr) *InequalityExpr {
	return &InequalityExpr{
		Left:  left,
		Right: right,
		Compare: func(lhs, rhs reflect.Value) bool {
			return reflectcmp.Compare(lhs, rhs) >= 0
		},
	}
}

func (e *InequalityExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	got := e.Compare(lhs, rhs)
	return reflect.ValueOf(got), nil
}

var _ Expr = (*InequalityExpr)(nil)
