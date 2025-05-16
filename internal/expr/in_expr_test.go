package expr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestIn(t *testing.T) {
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
			name:  "value present in slice",
			left:  exprtest.Integer(2),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(true),
		}, {
			name:  "value not present in slice",
			left:  exprtest.Integer(4),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(false),
		}, {
			name:  "empty slice",
			left:  exprtest.Integer(1),
			right: exprtest.Slice([]int{}),
			want:  reflect.ValueOf(false),
		}, {
			name:  "empty array (as slice)",
			left:  exprtest.Integer(1),
			right: exprtest.Slice([]int{}),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left is nil",
			left:    exprtest.Empty(),
			right:   exprtest.Slice([]int{1, 2, 3}),
			want:    reflect.Value{},
			wantErr: nil,
		}, {
			name:    "right is nil",
			left:    exprtest.Integer(1),
			right:   exprtest.Empty(),
			want:    reflect.Value{},
			wantErr: nil,
		}, {
			name:    "right is not slice or array",
			left:    exprtest.Integer(1),
			right:   exprtest.Integer(2),
			want:    reflect.Value{},
			wantErr: errs.ErrIncompatible,
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Slice([]int{1, 2, 3}),
			want:    reflect.Value{},
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			want:    reflect.Value{},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.In(tc.left, tc.right)

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

func TestNotInExpr(t *testing.T) {
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
			name:  "value present in slice",
			left:  exprtest.Integer(2),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(false),
		}, {
			name:  "value not present in slice",
			left:  exprtest.Integer(4),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(true),
		}, {
			name:  "empty slice",
			left:  exprtest.Integer(1),
			right: exprtest.Slice([]int{}),
			want:  reflect.ValueOf(true),
		}, {
			name:  "empty array (as slice)",
			left:  exprtest.Integer(1),
			right: exprtest.Slice([]int{}),
			want:  reflect.ValueOf(true),
		}, {
			name:    "left is nil",
			left:    exprtest.Empty(),
			right:   exprtest.Slice([]int{1, 2, 3}),
			want:    reflect.Value{},
			wantErr: nil,
		}, {
			name:    "right is nil",
			left:    exprtest.Integer(1),
			right:   exprtest.Empty(),
			want:    reflect.Value{},
			wantErr: nil,
		}, {
			name:    "right is not slice or array",
			left:    exprtest.Integer(1),
			right:   exprtest.Integer(2),
			want:    reflect.Value{},
			wantErr: errs.ErrIncompatible,
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Slice([]int{1, 2, 3}),
			want:    reflect.Value{},
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			want:    reflect.Value{},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.NotIn(tc.left, tc.right)

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
