package dcell

import (
	"encoding"
	"reflect"

	"rodusek.dev/pkg/dcell/internal/compile"
	"rodusek.dev/pkg/dcell/internal/expr"
)

// Option is an option that can be used to configure the dcell compiler.
type Option interface {
	apply(*compile.Config) error
}

type option func(*compile.Config) error

func (o option) apply(c *compile.Config) error {
	return o(c)
}

var _ Option = (*option)(nil)

// WithFunc adds a function to the dcell function table.
//
// fn may be any function type, including variadic functions, and must return
// (T, error) or just T. The implementation will automatically handle conversions
// from the input types to the function's parameter types and from the function's
// return types to the output types.
//
// Example:
//
//	dcell.WithFunc(func(base, exponent int) (int, error) {
//	    return int(math.Pow(float64(base), float64(exponent))), nil
//	})
func WithFunc(name string, fn any) Option {
	return option(func(c *compile.Config) error {
		return c.FuncTable.AddFunc(name, fn)
	})
}

// Expr is a compiled dcell expression that can be evaluated.
type Expr struct {
	expr    expr.Expr
	display string
}

// Compile compiles a dcell expression string into an Expr.
func Compile(expression string, opts ...Option) (*Expr, error) {
	cfg := &compile.Config{
		FuncTable: tableV1(),
	}
	for _, opt := range opts {
		if err := opt.apply(cfg); err != nil {
			return nil, err
		}
	}
	tree, err := compile.NewTree(expression, cfg)
	if err != nil {
		return nil, err
	}
	result := &Expr{
		expr:    tree,
		display: expression,
	}
	return result, nil
}

// MustCompile compiles a dcell expression string into an Expr and panics
// if it fails.
func MustCompile(expr string, opts ...Option) *Expr {
	e, err := Compile(expr, opts...)
	if err != nil {
		panic(err)
	}
	return e
}

// Eval evaluates the expression with the provided value context.
func (e *Expr) Eval(v any) (*Result, error) {
	rv := reflect.ValueOf(v)
	ctx := expr.NewContext(rv)
	got, err := e.expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	result := &Result{
		inner: got,
	}
	return result, nil
}

// MustEval evaluates the expression with the provided value context and panics
// if it fails.
func (e *Expr) MustEval(v any) *Result {
	result, err := e.Eval(v)
	if err != nil {
		panic(err)
	}
	return result
}

// String returns the string representation of the expression.
func (e *Expr) String() string {
	if e == nil {
		return ""
	}
	return e.display
}

// UnmarshalText unmarshals a dcell expression from text.
func (e *Expr) UnmarshalText(b []byte) error {
	exp, err := Compile(string(b))
	if err != nil {
		return err
	}
	*e = *exp
	return nil
}

var _ encoding.TextUnmarshaler = (*Expr)(nil)

// MarshalText marshals the expression to a textual format.
func (e *Expr) MarshalText() ([]byte, error) {
	return []byte(e.display), nil
}

var _ encoding.TextMarshaler = (*Expr)(nil)
