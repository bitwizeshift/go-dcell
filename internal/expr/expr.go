package expr

import (
	"reflect"
)

type Expr interface {
	Eval(ctx *Context) (reflect.Value, error)
}
