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

func TestImplies(t *testing.T) {
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
			name:  "both true",
			left:  exprtest.Boolean(true),
			right: exprtest.Boolean(true),
			want:  reflect.ValueOf(true),
		}, {
			name:  "true implies false",
			left:  exprtest.Boolean(true),
			right: exprtest.Boolean(false),
			want:  reflect.ValueOf(false),
		}, {
			name:  "both false",
			left:  exprtest.Boolean(false),
			right: exprtest.Boolean(false),
			want:  reflect.ValueOf(true),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Boolean(true),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Boolean(true),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:  "left is true, right is nil",
			left:  exprtest.Boolean(true),
			right: exprtest.Empty(),
			want:  reflect.ValueOf(false),
		}, {
			name:  "both nil",
			left:  exprtest.Empty(),
			right: exprtest.Empty(),
			want:  reflect.ValueOf(true),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Implies(tc.left, tc.right)

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
