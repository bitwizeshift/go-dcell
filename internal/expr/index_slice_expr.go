package expr

import (
	"fmt"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

type IndexSliceExpr struct {
	Begin, End Expr
}

func IndexSlice(begin, end Expr) *IndexSliceExpr {
	return &IndexSliceExpr{
		Begin: begin,
		End:   end,
	}
}

func (e *IndexSliceExpr) Eval(ctx *Context) (reflect.Value, error) {
	rv := ctx.Current
	if reflectconv.IsNil(rv) {
		return reflect.Value{}, nil
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return reflect.Value{}, fmt.Errorf("index slice: unable to slice %s", rv.Type().Name())
	}
	low, err := e.Begin.Eval(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	begin, err := reflectconv.Int(low)
	if err != nil {
		return reflect.Value{}, err
	}
	end := rv.Len()
	if e.End != nil {
		high, err := e.End.Eval(ctx)
		if err != nil {
			return reflect.Value{}, err
		}
		end, err = reflectconv.Int(high)
		if err != nil {
			return reflect.Value{}, err
		}
	}
	if begin < 0 {
		return reflect.Value{}, fmt.Errorf("begin index %d out of bounds", begin)
	}
	endIndex := end
	if end < 0 {
		endIndex = rv.Len() + end
	}
	if endIndex < 0 || endIndex > rv.Len() {
		return reflect.Value{}, fmt.Errorf("end index %d out of bounds", end)
	}
	if begin > endIndex {
		return reflect.Value{}, fmt.Errorf("begin index %d is greater than end index %d", begin, endIndex)
	}
	return rv.Slice(begin, endIndex), nil
}

var _ Expr = (*IndexSliceExpr)(nil)
