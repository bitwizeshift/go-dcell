package expr_test

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestBitwiseNot(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name:  "Bitwise not of zero",
			input: exprtest.Integer(uint64(0)),
			want:  reflect.ValueOf(^uint64(0)),
		}, {
			name:  "Bitwise not of one",
			input: exprtest.Integer(uint64(1)),
			want:  reflect.ValueOf(^uint64(1)),
		}, {
			name:  "Bitwise not of max uint64",
			input: exprtest.Integer(uint64(math.MaxUint64)),
			want:  reflect.ValueOf(^uint64(math.MaxUint64)),
		}, {
			name:    "Operand returns error",
			input:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Operand is not a number",
			input:   exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.BitwiseNot(tc.input)

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
