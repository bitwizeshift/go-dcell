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

func TestMultiplyExpr_Eval(t *testing.T) {
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
			name:  "multiply integers",
			left:  exprtest.Integer(6),
			right: exprtest.Integer(7),
			want:  reflect.ValueOf(int64(42)),
		}, {
			name:  "multiply floats",
			left:  exprtest.Float(2.5),
			right: exprtest.Float(4.0),
			want:  reflect.ValueOf(10.0),
		}, {
			name:  "multiply int and float",
			left:  exprtest.Integer(3),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(6.0),
		}, {
			name:  "multiply float and int",
			left:  exprtest.Float(3.0),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(6.0),
		}, {
			name:  "multiply by zero",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(42),
			want:  reflect.ValueOf(int64(0)),
		}, {
			name:  "multiply negative numbers",
			left:  exprtest.Integer(-3),
			right: exprtest.Integer(-7),
			want:  reflect.ValueOf(int64(21)),
		}, {
			name:    "overflow int64",
			left:    exprtest.Integer(math.MaxInt64),
			right:   exprtest.Integer(2),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "type error",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(2),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "one value is float, other is non-numeric",
			left:    exprtest.Float(2.0),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Multiply(tc.left, tc.right)

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

func TestDivideExpr_Eval(t *testing.T) {
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
			name:  "divide integers",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(7),
			want:  reflect.ValueOf(int64(6)),
		}, {
			name:  "divide floats",
			left:  exprtest.Float(10.0),
			right: exprtest.Float(2.5),
			want:  reflect.ValueOf(4.0),
		}, {
			name:  "divide int and float",
			left:  exprtest.Integer(6),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(3.0),
		}, {
			name:  "divide float and int",
			left:  exprtest.Float(6.0),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(3.0),
		}, {
			name:    "divide by zero int",
			left:    exprtest.Integer(42),
			right:   exprtest.Integer(0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "divide by zero float",
			left:    exprtest.Float(42.0),
			right:   exprtest.Float(0.0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "type error",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(2),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "one value is float, other is non-numeric",
			left:    exprtest.Float(2.0),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Divide(tc.left, tc.right)

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

func TestFloorDivideExpr_Eval(t *testing.T) {
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
			name:  "floor divide integers",
			left:  exprtest.Integer(7),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(int64(3)),
		}, {
			name:  "floor divide floats",
			left:  exprtest.Float(7.5),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(math.Floor(7.5 / 2.0)),
		}, {
			name:  "floor divide int and float",
			left:  exprtest.Integer(7),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(math.Floor(7.0 / 2.0)),
		}, {
			name:  "floor divide float and int",
			left:  exprtest.Float(7.5),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(math.Floor(7.5 / 2.0)),
		}, {
			name:    "floor divide by zero int",
			left:    exprtest.Integer(7),
			right:   exprtest.Integer(0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "floor divide by zero float",
			left:    exprtest.Float(7.5),
			right:   exprtest.Float(0.0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "type error",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(2),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "one value is float, other is non-numeric",
			left:    exprtest.Float(2.0),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.FloorDivide(tc.left, tc.right)

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

func TestModulusExpr_Eval(t *testing.T) {
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
			name:  "modulus integers",
			left:  exprtest.Integer(7),
			right: exprtest.Integer(3),
			want:  reflect.ValueOf(int64(1)),
		}, {
			name:  "modulus floats",
			left:  exprtest.Float(7.5),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(math.Mod(7.5, 2.0)),
		}, {
			name:  "modulus int and float",
			left:  exprtest.Integer(7),
			right: exprtest.Float(2.0),
			want:  reflect.ValueOf(math.Mod(7.0, 2.0)),
		}, {
			name:  "modulus float and int",
			left:  exprtest.Float(7.5),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(math.Mod(7.5, 2.0)),
		}, {
			name:    "modulus by zero int",
			left:    exprtest.Integer(7),
			right:   exprtest.Integer(0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "modulus by zero float",
			left:    exprtest.Float(7.5),
			right:   exprtest.Float(0.0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "type error",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(2),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "one value is float, other is non-numeric",
			left:    exprtest.Float(2.0),
			right:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Modulus(tc.left, tc.right)

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
