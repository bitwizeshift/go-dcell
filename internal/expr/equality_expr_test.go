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

func TestEquals(t *testing.T) {
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
			name:  "Equal integers",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(42),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Not equal integers",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Equal strings",
			left:  exprtest.String("foo"),
			right: exprtest.String("foo"),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Not equal strings",
			left:  exprtest.String("foo"),
			right: exprtest.String("bar"),
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Equal slices",
			left:  exprtest.Slice(1, 2, 3),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Not equal slices",
			left:  exprtest.Slice(1, 2, 3),
			right: exprtest.Slice(3, 2, 1),
			want:  reflect.ValueOf(false),
		},
		{
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		},
		{
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sut := expr.Equal(tc.left, tc.right)

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

func TestNotEquals(t *testing.T) {
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
			name:  "Not equal integers",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Equal integers",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(42),
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Not equal strings",
			left:  exprtest.String("foo"),
			right: exprtest.String("bar"),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Equal strings",
			left:  exprtest.String("foo"),
			right: exprtest.String("foo"),
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Not equal slices",
			left:  exprtest.Slice(1, 2, 3),
			right: exprtest.Slice(3, 2, 1),
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Equal slices",
			left:  exprtest.Slice(1, 2, 3),
			right: exprtest.Slice(1, 2, 3),
			want:  reflect.ValueOf(false),
		},
		{
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		},
		{
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sut := expr.NotEqual(tc.left, tc.right)

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
