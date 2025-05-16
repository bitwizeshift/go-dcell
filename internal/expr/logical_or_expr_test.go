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

func TestLogicalOr(t *testing.T) {
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
			name:  "left is true, right not evaluated",
			left:  exprtest.Boolean(true),
			right: exprtest.Error(errors.New("should not be called")),
			want:  reflect.ValueOf(true),
		}, {
			name:  "left is false, right is true",
			left:  exprtest.Boolean(false),
			right: exprtest.Boolean(true),
			want:  reflect.ValueOf(true),
		}, {
			name:  "left is false, right is false",
			left:  exprtest.Boolean(false),
			right: exprtest.Boolean(false),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Boolean(true),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Boolean(false),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:  "left is non-bool truthy (int)",
			left:  exprtest.Integer(42),
			right: exprtest.Boolean(false),
			want:  reflect.ValueOf(true),
		}, {
			name:  "left is non-bool falsy (zero int), right true",
			left:  exprtest.Integer(0),
			right: exprtest.Boolean(true),
			want:  reflect.ValueOf(true),
		}, {
			name:  "left and right are non-bool falsy",
			left:  exprtest.Integer(0),
			right: exprtest.String(""),
			want:  reflect.ValueOf(false),
		}, {
			name:  "left is nil, right is true",
			left:  exprtest.Empty(),
			right: exprtest.Boolean(true),
			want:  reflect.ValueOf(true),
		}, {
			name:  "left is nil, right is false",
			left:  exprtest.Empty(),
			right: exprtest.Boolean(false),
			want:  reflect.ValueOf(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.LogicalOr(tc.left, tc.right)

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
