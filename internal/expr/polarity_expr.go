package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/intconv"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type PolarityPlusExpr struct {
	Expr Expr
}

func PolarityPlus(e Expr) Expr {
	return PolarityPlusExpr{Expr: e}
}

func (e PolarityPlusExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv, err := e.Expr.Eval(ctx)
	if err != nil || reflectconv.IsNil(rv) {
		return reflect.Value{}, err
	}

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return rv, nil
	}
	return reflect.Value{}, fmt.Errorf("unary '+': not applicable to %s", rv.Type().Name())
}

type PolarityMinusExpr struct {
	Expr Expr
}

func PolarityMinus(e Expr) PolarityMinusExpr {
	return PolarityMinusExpr{
		Expr: e,
	}
}

func (e PolarityMinusExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv, err := e.Expr.Eval(ctx)
	if err != nil || reflectconv.IsNil(rv) {
		return reflect.Value{}, err
	}

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(-rv.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := intconv.Int64(rv.Uint())
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(-i), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(-rv.Float()), nil
	}
	return reflect.Value{}, fmt.Errorf("unary '-': not applicable to %v", rv.Type())
}

var _ Expr = (*PolarityMinusExpr)(nil)
