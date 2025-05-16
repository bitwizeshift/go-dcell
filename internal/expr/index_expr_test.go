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

func TestIndexExpr_Eval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name      string
		current   any
		indexExpr expr.Expr
		want      reflect.Value
		wantErr   error
	}{
		{
			name:      "index in slice",
			current:   []int{10, 20, 30},
			indexExpr: exprtest.Integer(1),
			want:      reflect.ValueOf(20),
		}, {
			name:      "index in array",
			current:   [3]int{1, 2, 3},
			indexExpr: exprtest.Integer(2),
			want:      reflect.ValueOf(3),
		}, {
			name:      "negative index",
			current:   []int{5, 6, 7},
			indexExpr: exprtest.Integer(-1),
			want:      reflect.ValueOf(7),
		}, {
			name:      "index out of bounds (positive)",
			current:   []int{1, 2},
			indexExpr: exprtest.Integer(5),
			wantErr:   cmpopts.AnyError,
		}, {
			name:      "index out of bounds (negative)",
			current:   []int{1, 2},
			indexExpr: exprtest.Integer(-3),
			wantErr:   cmpopts.AnyError,
		}, {
			name:      "not a slice or array",
			current:   42,
			indexExpr: exprtest.Integer(0),
			wantErr:   cmpopts.AnyError,
		}, {
			name:      "current is nil",
			current:   nil,
			indexExpr: exprtest.Integer(0),
			want:      reflect.Value{},
		}, {
			name:      "index expr returns error",
			current:   []int{1, 2, 3},
			indexExpr: exprtest.Error(testErr),
			wantErr:   testErr,
		}, {
			name:      "index expr not integer",
			current:   []int{1, 2, 3},
			indexExpr: exprtest.String("foo"),
			wantErr:   cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := &expr.Context{Current: reflect.ValueOf(tc.current)}
			sut := expr.Index(tc.indexExpr)

			got, err := sut.Eval(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}
