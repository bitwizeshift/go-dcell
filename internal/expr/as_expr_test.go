package expr_test

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestAsExpr(t *testing.T) {
	testErr := errors.New("test error")
	t.Parallel()
	testCases := []struct {
		name    string
		expr    expr.Expr
		as      expr.Type
		want    reflect.Value
		wantErr error
	}{
		{
			name: "Integer as Integer",
			expr: exprtest.Integer(42),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "Integer as Unsigned Integer",
			expr: exprtest.Integer(42),
			as:   expr.TypeUint,
			want: reflect.ValueOf(uint(42)),
		}, {
			name:    "Negative Integer as Unsigned Integer",
			expr:    exprtest.Integer(-42),
			as:      expr.TypeUint,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Integer as Float",
			expr: exprtest.Integer(42),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(42.0),
		}, {
			name: "Integer as String",
			expr: exprtest.Integer(42),
			as:   expr.TypeString,
			want: reflect.ValueOf("42"),
		}, {
			name: "Integer as Boolean true",
			expr: exprtest.Integer(1),
			as:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "Integer as Boolean false",
			expr: exprtest.Integer(0),
			as:   expr.TypeBool,
			want: reflect.ValueOf(false),
		}, {
			name: "Unsigned integer as Integer",
			expr: exprtest.Integer(uint(42)),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "Unsigned integer as Unsigned Integer",
			expr: exprtest.Integer(uint(42)),
			as:   expr.TypeUint,
			want: reflect.ValueOf(uint(42)),
		}, {
			name: "Unsigned integer as Float",
			expr: exprtest.Integer(uint(42)),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(42.0),
		}, {
			name: "Unsigned integer as String",
			expr: exprtest.Integer(uint(42)),
			as:   expr.TypeString,
			want: reflect.ValueOf("42"),
		}, {
			name: "Unsigned integer as Boolean true",
			expr: exprtest.Integer(uint(1)),
			as:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "Unsigned integer as Boolean false",
			expr: exprtest.Integer(uint(0)),
			as:   expr.TypeBool,
			want: reflect.ValueOf(false),
		}, {
			name: "Boolean false as Integer",
			expr: exprtest.Boolean(false),
			as:   expr.TypeInt,
			want: reflect.ValueOf(0),
		}, {
			name: "Boolean false as Unsigned Integer",
			expr: exprtest.Boolean(false),
			as:   expr.TypeUint,
			want: reflect.ValueOf(uint(0)),
		}, {
			name: "Boolean true as Integer",
			expr: exprtest.Boolean(true),
			as:   expr.TypeInt,
			want: reflect.ValueOf(1),
		}, {
			name: "Boolean true as Unsigned Integer",
			expr: exprtest.Boolean(true),
			as:   expr.TypeUint,
			want: reflect.ValueOf(uint(1)),
		}, {
			name: "Boolean false as Float",
			expr: exprtest.Boolean(false),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(0.0),
		}, {
			name: "Boolean true as Float",
			expr: exprtest.Boolean(true),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(1.0),
		}, {
			name: "Boolean false as String",
			expr: exprtest.Boolean(false),
			as:   expr.TypeString,
			want: reflect.ValueOf("false"),
		}, {
			name: "Boolean true as String",
			expr: exprtest.Boolean(true),
			as:   expr.TypeString,
			want: reflect.ValueOf("true"),
		}, {
			name: "Boolean as Boolean",
			expr: exprtest.Boolean(true),
			as:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "Float as Integer",
			expr: exprtest.Float(42.0),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "Float as Unsigned Integer",
			expr: exprtest.Float(42.0),
			as:   expr.TypeUint,
			want: reflect.ValueOf(uint(42)),
		}, {
			name: "Float as Float",
			expr: exprtest.Float(42.0),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(42.0),
		}, {
			name: "Float as String",
			expr: exprtest.Float(42.0),
			as:   expr.TypeString,
			want: reflect.ValueOf("42"),
		}, {
			name: "Float as Boolean true",
			expr: exprtest.Float(1.0),
			as:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "Float as Boolean false",
			expr: exprtest.Float(0.0),
			as:   expr.TypeBool,
			want: reflect.ValueOf(false),
		}, {
			name: "Float larger than int max",
			expr: exprtest.Float(float64(math.MaxUint64)),
			as:   expr.TypeInt,
			want: reflect.ValueOf(uint64(math.MaxUint64)),
		}, {
			name:    "Float smaller than int min",
			expr:    exprtest.Float(-9e100),
			as:      expr.TypeInt,
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Float larger than int max",
			expr:    exprtest.Float(9e100),
			as:      expr.TypeInt,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Hex string as Integer",
			expr: exprtest.String("0x2A"),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "Octal string as Integer",
			expr: exprtest.String("052"),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "Decimal string as Integer",
			expr: exprtest.String("42"),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name:    "Invalid string as integer",
			expr:    exprtest.String("invalid"),
			as:      expr.TypeInt,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Binary string as Integer",
			expr: exprtest.String("0b101010"),
			as:   expr.TypeInt,
			want: reflect.ValueOf(42),
		}, {
			name: "String as Float",
			expr: exprtest.String("42.0"),
			as:   expr.TypeFloat,
			want: reflect.ValueOf(42.0),
		}, {
			name:    "Invalid string as Float",
			expr:    exprtest.String("invalid"),
			as:      expr.TypeFloat,
			wantErr: cmpopts.AnyError,
		}, {
			name: "String as Boolean true",
			expr: exprtest.String("true"),
			as:   expr.TypeBool,
			want: reflect.ValueOf(true),
		}, {
			name: "String as Boolean false",
			expr: exprtest.String("false"),
			as:   expr.TypeBool,
			want: reflect.ValueOf(false),
		}, {
			name:    "Invalid string as Boolean",
			expr:    exprtest.String("invalid"),
			as:      expr.TypeBool,
			wantErr: cmpopts.AnyError,
		}, {
			name: "String as String",
			expr: exprtest.String("hello"),
			as:   expr.TypeString,
			want: reflect.ValueOf("hello"),
		}, {
			name: "Struct as Integer",
			expr: exprtest.Func(func(*expr.Context) (reflect.Value, error) {
				return reflect.ValueOf(struct{}{}), nil
			}),
			as:      expr.TypeInt,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Struct as Float",
			expr: exprtest.Func(func(*expr.Context) (reflect.Value, error) {
				return reflect.ValueOf(struct{}{}), nil
			}),
			as:      expr.TypeFloat,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Struct as String",
			expr: exprtest.Func(func(*expr.Context) (reflect.Value, error) {
				return reflect.ValueOf(struct{}{}), nil
			}),
			as:      expr.TypeString,
			wantErr: cmpopts.AnyError,
		}, {
			name: "Struct as Boolean",
			expr: exprtest.Func(func(*expr.Context) (reflect.Value, error) {
				return reflect.ValueOf(struct{}{}), nil
			}),
			as:      expr.TypeBool,
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Expr returns error",
			expr:    exprtest.Error(testErr),
			as:      expr.TypeInt,
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.As(tc.expr, tc.as)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AsExpr.Eval(...) error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("AsExpr.Eval(...) = %v, want %v", got, want)
			}
		})
	}
}
