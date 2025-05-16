package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// InExpr implements the `in` operator, which checks if the left operand is
// present in a slice or array represented by the right operand.
type InExpr struct {
	Left, Right Expr
	Transform   func(bool) bool
}

// In returns an [InExpr].
func In(left, right Expr) *InExpr {
	return &InExpr{
		Left:  left,
		Right: right,
		Transform: func(v bool) bool {
			return v
		},
	}
}

// NotIn returns an [InExpr] that negates the result of the `in` operator.
func NotIn(left, right Expr) Expr {
	return &InExpr{
		Left:  left,
		Right: right,
		Transform: func(v bool) bool {
			return !v
		},
	}
}

// Eval evaluates the InExpr.
func (e *InExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil || reflectconv.IsNil(lhs) || reflectconv.IsNil(rhs) {
		return reflect.Value{}, err
	}
	kind := rhs.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return reflect.Value{}, fmt.Errorf(
			"%w: right operand must be a slice or array, got %v",
			errs.ErrIncompatible,
			rhs.Type(),
		)
	}
	for i := range rhs.Len() {
		elem := rhs.Index(i)
		if reflectcmp.Equal(lhs, elem) {
			return reflect.ValueOf(e.Transform(true)), nil
		}
	}
	return reflect.ValueOf(e.Transform(false)), nil
}

var _ Expr = (*InExpr)(nil)
