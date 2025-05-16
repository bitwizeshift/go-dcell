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

func TestBitwiseAndEval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        reflect.Value
		wantErr     error
	}{
		{
			name:  "Both zero",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(uint64(0)),
		}, {
			name:  "Zero and positive",
			left:  exprtest.Integer(0),
			right: exprtest.Integer(42),
			want:  reflect.ValueOf(uint64(0)),
		}, {
			name:  "Positive and zero",
			left:  exprtest.Integer(42),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(uint64(0)),
		}, {
			name:  "All bits set and zero",
			left:  exprtest.Integer(math.MaxInt64),
			right: exprtest.Integer(0),
			want:  reflect.ValueOf(uint64(0)),
		}, {
			name:  "All bits set and all bits set",
			left:  exprtest.Integer(math.MaxInt64),
			right: exprtest.Integer(math.MaxInt64),
			want:  reflect.ValueOf(uint64(math.MaxInt64)),
		}, {
			name:  "Some bits overlap",
			left:  exprtest.Integer(0b1101),
			right: exprtest.Integer(0b1011),
			want:  reflect.ValueOf(uint64(0b1001)),
		}, {
			name:  "No bits overlap",
			left:  exprtest.Integer(0b1000),
			right: exprtest.Integer(0b0111),
			want:  reflect.ValueOf(uint64(0)),
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
			name:    "both return error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Error(testErr),
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
			sut := expr.BitwiseAnd(tc.left, tc.right)

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
