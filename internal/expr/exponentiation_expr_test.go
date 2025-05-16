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

func TestPower(t *testing.T) {
	t.Parallel()

	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		left    expr.Expr
		right   expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name:  "integer exponentiation",
			left:  exprtest.Integer(2),
			right: exprtest.Integer(3),
			want:  reflect.ValueOf(int64(8)),
		}, {
			name:  "zero exponent",
			left:  exprtest.Integer(5),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(int64(1)),
		}, {
			name:  "zero base, positive exponent",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(3),
			want:  reflect.ValueOf(int64(0)),
		}, {
			name:  "float exponentiation",
			left:  exprtest.Float(2.0),
			right: exprtest.Float(3.0),
			want:  reflect.ValueOf(math.Pow(2.0, 3.0)),
		}, {
			name:  "float base, integer exponent",
			left:  exprtest.Float(2.0),
			right: exprtest.Integer(4),
			want:  reflect.ValueOf(math.Pow(2.0, 4.0)),
		}, {
			name:  "integer base, float exponent",
			left:  exprtest.Integer(9),
			right: exprtest.Float(0.5),
			want:  reflect.ValueOf(math.Pow(9.0, 0.5)),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(2),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(2),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "one side is float, other is non-numeric",
			left:    exprtest.Float(2.0),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "one side is integer, other is non-numeric",
			left:    exprtest.Integer(2),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "integer power overflows",
			left:    exprtest.Integer(2),
			right:   exprtest.Integer(66),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Power(tc.left, tc.right)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("PowerExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("PowerExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
