package expr_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestLiteralExpr(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		input   any
		want    reflect.Value
		wantErr error
	}{
		{
			name:  "Input is empty",
			input: (any)(nil),
			want:  reflect.Value{},
		}, {
			name:  "Input is value",
			input: 42,
			want:  reflect.ValueOf(42),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Literal(tc.input)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("LiteralExpr(%v) = %v, want %v", tc.input, err, tc.wantErr)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("LiteralExpr(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}
