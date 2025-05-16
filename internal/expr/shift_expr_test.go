package expr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestBitwiseShiftLeft(t *testing.T) {
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
			name:  "Simple left shift",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(3),
			want:  reflect.ValueOf(uint64(8)),
		}, {
			name:  "Zero shift",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(uint64(42)),
		}, {
			name:  "Shift by 63 (max for uint64)",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(63),
			want:  reflect.ValueOf(uint64(1) << 63),
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "input is not a number",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.BitwiseShiftLeft(tc.left, tc.right)

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

func TestBitwiseShiftRight(t *testing.T) {
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
			name:  "Simple right shift",
			left:  exprtest.Integer(8),
			right: exprtest.Integer(3),
			want:  reflect.ValueOf(uint64(1)),
		}, {
			name:  "Zero shift",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(uint64(42)),
		}, {
			name:  "Shift by 63 (max for uint64)",
			left:  exprtest.Integer(int(^uint64(0) >> 1)), // largest int
			right: exprtest.Integer(63),
			want:  reflect.ValueOf(uint64(int(^uint64(0)>>1)) >> 63),
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "input is not a number",
			left:    exprtest.String("foo"),
			right:   exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.BitwiseShiftRight(tc.left, tc.right)

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
