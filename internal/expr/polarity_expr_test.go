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
	"rodusek.dev/pkg/dcell/internal/intconv"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestPolarityPlus(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name: "int value",
			expr: exprtest.Integer(42),
			want: reflect.ValueOf(int64(42)),
		}, {
			name: "float value",
			expr: exprtest.Float(3.14),
			want: reflect.ValueOf(3.14),
		}, {
			name: "uint value",
			expr: exprtest.Integer(uint64(7)),
			want: reflect.ValueOf(uint64(7)),
		}, {
			name:    "error",
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name: "nil value",
			expr: exprtest.Empty(),
			want: reflect.Value{},
		}, {
			name:    "not applicable type (string)",
			expr:    exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.PolarityPlus(tc.expr)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestPolarityMinus(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name: "int value",
			expr: exprtest.Integer(42),
			want: reflect.ValueOf(int64(-42)),
		}, {
			name: "negative int value",
			expr: exprtest.Integer(-7),
			want: reflect.ValueOf(int64(7)),
		}, {
			name: "float value",
			expr: exprtest.Float(3.14),
			want: reflect.ValueOf(-3.14),
		}, {
			name: "negative float value",
			expr: exprtest.Float(-2.5),
			want: reflect.ValueOf(2.5),
		}, {
			name: "uint value",
			expr: exprtest.Integer(uint64(7)),
			want: reflect.ValueOf(int64(-7)),
		}, {
			name:    "error",
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name: "nil value",
			expr: exprtest.Empty(),
			want: reflect.Value{},
		}, {
			name:    "not applicable type (string)",
			expr:    exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "integer overflows",
			expr:    exprtest.Integer(uint64(math.MaxUint64)),
			wantErr: intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.PolarityMinus(tc.expr)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}
