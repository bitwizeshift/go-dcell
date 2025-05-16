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

func sequence(exprs ...expr.Expr) expr.SequenceExpr {
	if len(exprs) == 0 {
		return expr.SequenceExpr{}
	}
	return expr.Sequence(exprs[0], exprs[1:]...)
}

func TestSequenceExpr_Eval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		exprs   []expr.Expr
		current reflect.Value
		want    reflect.Value
		wantErr error
	}{
		{
			name:    "empty sequence returns current",
			exprs:   []expr.Expr{},
			current: reflect.ValueOf(42),
			want:    reflect.ValueOf(42),
		}, {
			name:    "single expression returns its value",
			exprs:   []expr.Expr{exprtest.Integer(7)},
			current: reflect.ValueOf(0),
			want:    reflect.ValueOf(7),
		}, {
			name:    "multiple expressions chain values",
			exprs:   []expr.Expr{exprtest.Integer(1), exprtest.Integer(2), exprtest.Integer(3)},
			current: reflect.ValueOf(0),
			want:    reflect.ValueOf(3),
		}, {
			name:    "expression returns error propagates",
			exprs:   []expr.Expr{exprtest.Integer(1), exprtest.Error(testErr), exprtest.Integer(3)},
			current: reflect.ValueOf(0),
			wantErr: testErr,
		}, {
			name:    "expression returns nil stops sequence",
			exprs:   []expr.Expr{exprtest.Integer(1), exprtest.Empty(), exprtest.Integer(3)},
			current: reflect.ValueOf(0),
			wantErr: nil,
		}, {
			name: "first expression is sequence itself",
			exprs: []expr.Expr{
				expr.SequenceExpr{exprtest.Integer(1)},
				exprtest.Integer(2),
			},
			current: reflect.ValueOf(0),
			want:    reflect.ValueOf(2),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := sequence(tc.exprs...)
			ctx := &expr.Context{Current: tc.current}

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
