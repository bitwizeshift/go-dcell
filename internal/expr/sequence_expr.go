package expr

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// SequenceExpr is a utility expression for forming multiple expressions into
// a sequence, passing the result of the previous expression to the next one.
type SequenceExpr []Expr

// Sequence creates a new SequenceExpr with the given expressions. The first
// expression is the initial value, and the rest are the subsequent expressions
// to be evaluated in order.
//
// If the first expression is already a SequenceExpr, it will be appended to
// the rest of the expressions.
func Sequence(first Expr, rest ...Expr) SequenceExpr {
	if seq, ok := first.(SequenceExpr); ok {
		return append(seq, rest...)
	}
	return append(SequenceExpr{first}, rest...)
}

// Eval evaluates the sequence of expressions in order, passing the result of
// each expression to the next one. If any expression returns an error or a nil
// value, the evaluation stops and returns the error or nil value.
func (e SequenceExpr) Eval(ctx *Context) (reflect.Value, error) {
	current := ctx.Current
	for _, expr := range e {
		result, err := expr.Eval(ctx.Next(current))
		if err != nil || reflectconv.IsNil(result) {
			return reflect.Value{}, err
		}
		current = result
	}
	return current, nil
}
