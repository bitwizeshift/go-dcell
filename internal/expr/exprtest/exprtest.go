package exprtest

import (
	"reflect"

	"golang.org/x/exp/constraints"
	"rodusek.dev/pkg/dcell/internal/expr"
)

// Func is a function type that implements the expr.Expr interface.
type Func func(*expr.Context) (reflect.Value, error)

// Eval evaluates the function and returns the result.
func (f Func) Eval(ctx *expr.Context) (reflect.Value, error) {
	return f(ctx)
}

var _ expr.Expr = (*Func)(nil)

// Boolean creates an [Expr] that returns a boolean value.
func Boolean[T ~bool](b T) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.ValueOf(b), nil
	})
}

// Integer creates an [Expr] that returns an integer value.
func Integer[T constraints.Integer](i T) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.ValueOf(i), nil
	})
}

// Float creates an [Expr] that returns a float value.
func Float[T constraints.Float](f T) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.ValueOf(f), nil
	})
}

// String creates an [Expr] that returns a string value.
func String[T ~string](s T) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.ValueOf(s), nil
	})
}

// Slice creates an [Expr] that returns a slice of values.
func Slice[T any](s ...T) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.ValueOf(s), nil
	})
}

// Error creates an [Expr] that returns an error.
func Error(err error) expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.Value{}, err
	})
}

// Empty creates an [Expr] that returns a nil value.
func Empty() expr.Expr {
	return Func(func(*expr.Context) (reflect.Value, error) {
		return reflect.Value{}, nil
	})
}
