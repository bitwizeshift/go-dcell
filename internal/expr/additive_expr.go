package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type AddExpr struct {
	Left, Right Expr
}

func Add(left, right Expr) *AddExpr {
	return &AddExpr{
		Left:  left,
		Right: right,
	}
}

func (e *AddExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}

	if reflectconv.IsInt(lhs.Type()) && reflectconv.IsInt(rhs.Type()) {
		ints, err := reflectconv.Int64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsInt, rhsInt := ints[0], ints[1]
		return reflect.ValueOf(lhsInt + rhsInt), nil
	} else if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		return reflect.ValueOf(lhsFloat + rhsFloat), nil
	} else if reflectconv.IsString(lhs.Type()) && reflectconv.IsString(rhs.Type()) {
		result := lhs.String() + rhs.String()
		return reflect.ValueOf(result), nil
	}
	return reflect.Value{}, fmt.Errorf("%w: operation for %v + %v is undefined", errs.ErrIncompatible, lhs.Type(), rhs.Type())
}

var _ Expr = (*AddExpr)(nil)

type SubtractExpr struct {
	Left, Right Expr
}

func Subtract(left, right Expr) *SubtractExpr {
	return &SubtractExpr{
		Left:  left,
		Right: right,
	}
}

func (e *SubtractExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsInt(lhs.Type()) && reflectconv.IsInt(rhs.Type()) {
		ints, err := reflectconv.Int64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsInt, rhsInt := ints[0], ints[1]
		return reflect.ValueOf(lhsInt - rhsInt), nil
	} else if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		return reflect.ValueOf(lhsFloat - rhsFloat), nil
	}

	return reflect.Value{}, fmt.Errorf("%w: operation for %v - %v is undefined", errs.ErrIncompatible, lhs.Type(), rhs.Type())
}

var _ Expr = (*AddExpr)(nil)
