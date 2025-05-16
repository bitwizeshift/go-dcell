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

func TestBitwiseOrExpr_Eval(t *testing.T) {
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
			name:  "Bitwise or of two positive integers",
			left:  exprtest.Integer(6),       // 110
			right: exprtest.Integer(3),       // 011
			want:  reflect.ValueOf(int64(7)), // 111
		}, {
			name:  "Bitwise or with zero",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(5),
			want:  reflect.ValueOf(int64(5)),
		}, {
			name:  "Bitwise or of two zeros",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(int64(0)),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "both return error (left prioritized)",
			left:    exprtest.Error(testErr),
			right:   exprtest.Error(errors.New("other error")),
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
			sut := expr.BitwiseOr(tc.left, tc.right)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; tc.wantErr == nil && !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestBitwiseXorExpr_Eval(t *testing.T) {
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
			name:  "Bitwise xor of two positive integers",
			left:  exprtest.Integer(6),       // 110
			right: exprtest.Integer(3),       // 011
			want:  reflect.ValueOf(int64(5)), // 101
		}, {
			name:  "Bitwise xor with zero",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(5),
			want:  reflect.ValueOf(int64(5)),
		}, {
			name:  "Bitwise xor of two zeros",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(int64(0)),
		}, {
			name:    "Bitwise xor of negative and positive",
			left:    exprtest.Integer(-2),
			right:   exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "both return error (left prioritized)",
			left:    exprtest.Error(testErr),
			right:   exprtest.Error(errors.New("other error")),
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
			sut := expr.BitwiseXor(tc.left, tc.right)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; tc.wantErr == nil && !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}
