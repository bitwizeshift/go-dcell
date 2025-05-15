package expr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
)

func TestLogicalNotExpr(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name:  "Input is empty",
			input: exprtest.Empty(),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Input is falsey",
			input: exprtest.Boolean(false),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Input is truthy",
			input: exprtest.Integer(1),
			want:  reflect.ValueOf(false),
		}, {
			name:    "Input returns error",
			input:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.LogicalNot(tc.input)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("LogicalNotExpr.Eval() error = %v, want %v", got, want)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LogicalNotExpr.Eval() got = %v, want %v", got, tc.want)
			}
		})
	}
}
