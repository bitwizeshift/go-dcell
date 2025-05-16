package expr

import (
	"fmt"
	"math"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// PowerExpr implements the exponentiation operator (**).
type PowerExpr struct {
	Left, Right Expr
}

// Power creates a new PowerExpr with the given left and right operands.
func Power(left, right Expr) *PowerExpr {
	return &PowerExpr{
		Left:  left,
		Right: right,
	}
}

// Eval evaluates the PowerExpr.
func (e *PowerExpr) Eval(ctx *Context) (reflect.Value, error) {
	lhs, rhs, err := evalTwo(ctx, e.Left, e.Right)
	if err != nil {
		return reflect.Value{}, err
	}
	if reflectconv.IsFloat(lhs.Type()) || reflectconv.IsFloat(rhs.Type()) {
		fs, err := reflectconv.Float64s(lhs, rhs)
		if err != nil {
			return reflect.Value{}, err
		}
		lhsFloat, rhsFloat := fs[0], fs[1]
		return reflect.ValueOf(math.Pow(lhsFloat, rhsFloat)), nil
	}
	is, err := reflectconv.Int64s(lhs, rhs)
	if err != nil {
		return reflect.Value{}, err
	}
	lhsInt, rhsInt := is[0], is[1]

	// This is not a perfect check for overflow, since large values of float64
	// will lose precision -- missing some of the overflow check. This is just
	// a best-effort, since the "proper" way is to check at each stage of the
	// multiplication in intPow, and that would be a large pessimization.
	if pow := math.Pow(float64(lhsInt), float64(rhsInt)); pow > math.MaxInt64 {
		return reflect.Value{}, fmt.Errorf("result of %d^%d overflows integer", lhsInt, rhsInt)
	}
	return reflect.ValueOf(intPow(lhsInt, rhsInt)), nil
}

func intPow(base, exp int64) int64 {
	if base == 0 {
		return 0
	}
	if exp == 0 {
		return 1
	}
	result := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

var _ Expr = (*PowerExpr)(nil)
