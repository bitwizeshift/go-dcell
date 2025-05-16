package expr

import (
	"reflect"
)

type fn = func(args ...reflect.Value) (reflect.Value, error)

type FreeFuncExpr struct {
	Args []Expr
	Fn   fn
}

func FreeFunc(fn fn, args ...Expr) *FreeFuncExpr {
	return &FreeFuncExpr{
		Args: args,
		Fn:   fn,
	}
}

func (f *FreeFuncExpr) Eval(ctx *Context) (reflect.Value, error) {
	args := make([]reflect.Value, 0, len(f.Args))
	for _, arg := range f.Args {
		result, err := arg.Eval(ctx)
		if err != nil {
			return reflect.Value{}, err
		}
		args = append(args, result)
	}
	got, err := f.Fn(args...)
	if err != nil {
		return reflect.Value{}, err
	}
	return got, nil
}

var _ Expr = (*FreeFuncExpr)(nil)

type MemberFuncExpr struct {
	Args []Expr
	Fn   fn
}

func MemberFunc(fn fn, args ...Expr) *MemberFuncExpr {
	return &MemberFuncExpr{
		Args: args,
		Fn:   fn,
	}
}

func (e *MemberFuncExpr) Eval(ctx *Context) (reflect.Value, error) {
	current := ctx.Current
	if !current.IsValid() {
		return reflect.Value{}, nil
	}
	args := make([]reflect.Value, 0, len(e.Args)+1)
	args = append(args, current)
	for _, arg := range e.Args {
		result, err := arg.Eval(ctx)
		if err != nil {
			return reflect.Value{}, err
		}
		args = append(args, result)
	}
	return e.Fn(args...)
}
