package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// IsExpr implements the 'is' operator for checking the type of a value.
type IsExpr struct {
	Expr Expr
	Type Type
}

// Is creates an [IsExpr].
func Is(ex Expr, ty Type) *IsExpr {
	return &IsExpr{
		Expr: ex,
		Type: ty,
	}
}

// IsNot creates an [Expr] that checks if the value is _not_ of the specified type.
func IsNot(ex Expr, ty Type) Expr {
	return LogicalNot(Is(ex, ty))
}

// Eval evaluates the 'is' expression.
func (e *IsExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv, err := e.Expr.Eval(ctx)
	if err != nil || reflectconv.IsNil(rv) {
		return reflect.Value{}, err
	}
	switch rv.Kind() {
	case reflect.String:
		return reflect.ValueOf(e.Type == TypeString), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(e.Type == TypeInt), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflect.ValueOf(e.Type == TypeUint), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(e.Type == TypeFloat), nil
	case reflect.Bool:
		return reflect.ValueOf(e.Type == TypeBool), nil
	}
	return reflect.ValueOf(false), nil
}

var _ Expr = (*IsExpr)(nil)
