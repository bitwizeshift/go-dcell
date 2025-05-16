package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type IndexExpr struct {
	Index Expr
}

func Index(ex Expr) IndexExpr {
	return IndexExpr{
		Index: ex,
	}
}

func (e IndexExpr) Eval(ctx *Context) (reflect.Value, error) {
	_ = ctx
	current := ctx.Current
	if reflectconv.IsNil(current) {
		return reflect.Value{}, nil
	}

	rv := reflectconv.Deref(current)
	got, err := e.Index.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	index, err := reflectconv.Int(got)
	if err != nil {
		return reflect.Value{}, err
	}

	rt := rv.Type()
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return reflect.Value{}, fmt.Errorf("index %d does not exist in %s", index, rt.Name())
	}
	i := index
	if i < 0 {
		i = rv.Len() + i
	}

	if i < 0 || i >= rv.Len() {
		return reflect.Value{}, fmt.Errorf("index %d out of bounds for %s", index, rt.Name())
	}

	rfield := rv.Index(i)
	return rfield, nil
}

var _ Expr = (*IndexExpr)(nil)
