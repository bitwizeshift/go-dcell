package expr_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestIsExpr(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		expr    expr.Expr
		ty      expr.Type
		want    reflect.Value
		wantErr error
	}{
		{
			name: "Input is string, type is string",
			expr: exprtest.String("foo"),
			ty:   expr.TypeString,
			want: reflect.ValueOf(true),
		}, {
			name: "Input is string, type is not string",
			expr: exprtest.String("foo"),
			ty:   expr.TypeInt,
			want: reflect.ValueOf(false),
		}, {
			name: "Input is int, type is int",
			expr: exprtest.Integer(42),
			ty:   expr.TypeInt,
			want: reflect.ValueOf(true),
		}, {
			name: "Input is int, type is not int",
			expr: exprtest.Integer(42),
			ty:   expr.TypeString,
			want: reflect.ValueOf(false),
		}, {
			name: "input is uint, type is uint",
			expr: exprtest.Integer(uint(42)),
			ty:   expr.TypeUint,
			want: reflect.ValueOf(true),
		}, {
			name: "input is uint, type is not uint",
			expr: exprtest.Integer(uint(42)),
			ty:   expr.TypeInt,
			want: reflect.ValueOf(false),
		}, {
			name: "input is float, type is float",
			expr: exprtest.Float(42.0),
			ty:   expr.TypeFloat,
			want: reflect.ValueOf(true),
		}, {
			name: "input is float, type is not float",
			expr: exprtest.Float(42.0),
			ty:   expr.TypeInt,
			want: reflect.ValueOf(false),
		}, {
			name: "input is bool, type is bool",
			expr: exprtest.Boolean(true),
			ty:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "input is bool, type is not bool",
			expr: exprtest.Boolean(true),
			ty:   expr.TypeString,
			want: reflect.ValueOf(false),
		}, {
			name: "input is nil, type is string",
			expr: exprtest.Empty(),
			ty:   expr.TypeString,
		}, {
			name: "input is struct, type is not struct",
			expr: exprtest.Func(func(*expr.Context) (reflect.Value, error) {
				return reflect.ValueOf(struct{}{}), nil
			}),
			ty:   expr.TypeString,
			want: reflect.ValueOf(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Is(tc.expr, tc.ty)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsExpr() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("IsExpr(%v, %v) = %v, want %v", tc.expr, tc.ty, got, want)
			}
		})
	}
}
