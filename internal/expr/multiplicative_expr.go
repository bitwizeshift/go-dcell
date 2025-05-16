package expr

import (
	"fmt"
	"math"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type MultiplyExpr struct {
	Left, Right Expr
}

func Multiply(left, right Expr) *MultiplyExpr {
	return &MultiplyExpr{
		Left:  left,
		Right: right,
	}
}

func (e *MultiplyExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		return reflect.ValueOf(lhsFloat * rhsFloat), nil
	}
	ints, err := reflectconv.Int64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := ints[0], ints[1]
	if product := float64(lhsInt) * float64(rhsInt); product > float64(math.MaxInt64) {
		return reflect.Value{}, fmt.Errorf("product of %d and %d is too large", lhsInt, rhsInt)
	}
	return reflect.ValueOf(lhsInt * rhsInt), nil
}

var _ Expr = (*MultiplyExpr)(nil)

type DivideExpr struct {
	Left, Right Expr
}

func Divide(left, right Expr) *DivideExpr {
	return &DivideExpr{
		Left:  left,
		Right: right,
	}
}

func (e *DivideExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		if rhsFloat == 0 {
			return reflect.Value{}, fmt.Errorf("division by zero")
		}
		return reflect.ValueOf(lhsFloat / rhsFloat), nil
	}
	ints, err := reflectconv.Int64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := ints[0], ints[1]
	if rhsInt == 0 {
		return reflect.Value{}, fmt.Errorf("division by zero")
	}
	return reflect.ValueOf(lhsInt / rhsInt), nil
}

var _ Expr = (*DivideExpr)(nil)

type FloorDivideExpr struct {
	Left, Right Expr
}

func FloorDivide(left, right Expr) *FloorDivideExpr {
	return &FloorDivideExpr{
		Left:  left,
		Right: right,
	}
}

func (e *FloorDivideExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		if rhsFloat == 0 {
			return reflect.Value{}, fmt.Errorf("division by zero")
		}
		return reflect.ValueOf(math.Floor(lhsFloat / rhsFloat)), nil
	}
	ints, err := reflectconv.Int64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := ints[0], ints[1]
	if rhsInt == 0 {
		return reflect.Value{}, fmt.Errorf("division by zero")
	}
	return reflect.ValueOf(lhsInt / rhsInt), nil
}

var _ Expr = (*FloorDivideExpr)(nil)

// ModulusExpr represents a modulus expression.
type ModulusExpr struct {
	Left, Right Expr
}

// Modulus returns a [ModulusExpr] with the given left and right operands.
func Modulus(left, right Expr) *ModulusExpr {
	return &ModulusExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the ModulusExpr.
func (e *ModulusExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		floats, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := floats[0], floats[1]
		if rhsFloat == 0 {
			return reflect.Value{}, fmt.Errorf("modulo by zero")
		}
		return reflect.ValueOf(math.Mod(lhsFloat, rhsFloat)), nil
	}
	ints, err := reflectconv.Int64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := ints[0], ints[1]
	if rhsInt == 0 {
		return reflect.Value{}, fmt.Errorf("modulo by zero")
	}
	return reflect.ValueOf(lhsInt % rhsInt), nil
}

var _ Expr = (*ModulusExpr)(nil)
