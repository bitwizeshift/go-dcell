package expr

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"rodusek.dev/pkg/dcell/internal/intconv"
)

// AsExpr is an expression that converts a value to a different type.
type AsExpr struct {
	Expr Expr
	Type Type
}

// As creates a new AsExpr with the given expression and type.
func As(expr Expr, ty Type) *AsExpr {
	return &AsExpr{
		Expr: expr,
		Type: ty,
	}
}

// Eval evaluates the AsExpr and converts the value to the specified type.
func (e *AsExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv, err := e.Expr.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	switch e.Type {
	case TypeInt:
		return e.asInt(rv)
	case TypeUint:
		return e.asUint(rv)
	case TypeFloat:
		return e.asFloat(rv)
	case TypeString:
		return e.asString(rv)
	case TypeBool:
		return e.asBool(rv)
	}

	// This should be unreachable
	return reflect.Value{}, nil
}

func (a *AsExpr) asInt(rv reflect.Value) (reflect.Value, error) {
	intTo := func(rv reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(rv.Int()), nil
	}
	uintTo := func(rv reflect.Value) (reflect.Value, error) {
		v, err := intconv.Int64(rv.Uint())
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(v), nil
	}
	return a.asIntImpl(rv, intTo, uintTo)
}

func (e *AsExpr) asUint(rv reflect.Value) (reflect.Value, error) {
	intTo := func(rv reflect.Value) (reflect.Value, error) {
		v, err := intconv.Uint64(rv.Int())
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(v), nil
	}
	uintTo := func(rv reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(rv.Uint()), nil
	}
	return e.asIntImpl(rv, intTo, uintTo)
}

func (e *AsExpr) asIntImpl(rv reflect.Value, intTo, uintTo func(reflect.Value) (reflect.Value, error)) (reflect.Value, error) {
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			return reflect.ValueOf(1), nil
		}
		return reflect.ValueOf(0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intTo(rv)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintTo(rv)
	case reflect.Float32, reflect.Float64:
		f := rv.Float()
		if f >= float64(math.MaxInt64) {
			if f <= float64(math.MaxUint64) {
				return reflect.ValueOf(uint64(f)), nil
			}
			return reflect.Value{}, fmt.Errorf("float value %f is too large to convert to int", f)
		}
		if f < float64(math.MinInt64) {
			return reflect.Value{}, fmt.Errorf("float value %f is too small to convert to int", f)
		}
		return reflect.ValueOf(int64(rv.Float())), nil
	case reflect.String:
		str := rv.String()
		var v int64
		var err error
		if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
			v, err = strconv.ParseInt(str[2:], 16, 64)
		} else if strings.HasPrefix(str, "0b") || strings.HasPrefix(str, "0B") {
			v, err = strconv.ParseInt(str[2:], 2, 64)
		} else if strings.HasPrefix(str, "0") {
			v, err = strconv.ParseInt(str, 8, 64)
		} else {
			v, err = strconv.ParseInt(str, 10, 64)
		}
		if err != nil {
			return reflect.Value{}, err
		}
		return intTo(reflect.ValueOf(v))
	}
	return reflect.Value{}, fmt.Errorf("cannot convert %s to int", rv.Type().Name())
}

func (e *AsExpr) asFloat(rv reflect.Value) (reflect.Value, error) {
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			return reflect.ValueOf(1.0), nil
		}
		return reflect.ValueOf(0.0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(float64(rv.Int())), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(float64(rv.Uint())), nil
	case reflect.Float32, reflect.Float64:
		return rv, nil
	case reflect.String:
		str := rv.String()
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(f), nil
	}
	return reflect.Value{}, fmt.Errorf("cannot convert %s to float", rv.Type().Name())
}

func (e *AsExpr) asString(rv reflect.Value) (reflect.Value, error) {
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			return reflect.ValueOf("true"), nil
		}
		return reflect.ValueOf("false"), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(strconv.FormatInt(rv.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(strconv.FormatUint(rv.Uint(), 10)), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(strconv.FormatFloat(rv.Float(), 'f', -1, 64)), nil
	case reflect.String:
		return rv, nil
	}
	return reflect.Value{}, fmt.Errorf("cannot convert %s to string", rv.Type().Name())
}

func (e *AsExpr) asBool(rv reflect.Value) (reflect.Value, error) {
	switch rv.Kind() {
	case reflect.Bool:
		return rv, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.Int() == 0 {
			return reflect.ValueOf(false), nil
		}
		return reflect.ValueOf(true), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if rv.Uint() == 0 {
			return reflect.ValueOf(false), nil
		}
		return reflect.ValueOf(true), nil
	case reflect.Float32, reflect.Float64:
		if rv.Float() == 0.0 {
			return reflect.ValueOf(false), nil
		}
		return reflect.ValueOf(true), nil
	case reflect.String:
		str := rv.String()
		switch str {
		case "0", "false":
			return reflect.ValueOf(false), nil
		case "1", "true":
			return reflect.ValueOf(true), nil
		}
		return reflect.Value{}, fmt.Errorf("cannot convert string %q to bool", str)
	}
	return reflect.Value{}, fmt.Errorf("cannot convert %s to bool", rv.Type().Name())
}

var _ Expr = (*AsExpr)(nil)
