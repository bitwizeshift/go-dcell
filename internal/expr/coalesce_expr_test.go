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

func TestCoalesceExpr_Eval(t *testing.T) {
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
			name:  "Left non-nil, right not evaluated",
			left:  exprtest.Integer(42),
			right: exprtest.Error(errors.New("should not be called")),
			want:  reflect.ValueOf(42),
		}, {
			name:  "Left nil, right non-nil",
			left:  exprtest.Empty(),
			right: exprtest.String("foo"),
			want:  reflect.ValueOf("foo"),
		}, {
			name:  "Both nil",
			left:  exprtest.Empty(),
			right: exprtest.Empty(),
			want:  reflect.Value{},
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "Left nil, right returns error",
			left:    exprtest.Empty(),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Both return error (left error short-circuits)",
			left:    exprtest.Error(testErr),
			right:   exprtest.Error(errors.New("should not be called")),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Coalesce(tc.left, tc.right)

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
