package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
)

func TestType_UnmarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   string
		want    expr.Type
		wantErr error
	}{
		{
			name:  "string",
			input: "string",
			want:  expr.TypeString,
		}, {
			name:  "int",
			input: "int",
			want:  expr.TypeInt,
		}, {
			name:  "uint",
			input: "uint",
			want:  expr.TypeUint,
		}, {
			name:  "float",
			input: "float",
			want:  expr.TypeFloat,
		}, {
			name:  "bool",
			input: "bool",
			want:  expr.TypeBool,
		}, {
			name:    "invalid type",
			input:   "invalid",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var out expr.Type
			err := out.UnmarshalText([]byte(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("UnmarshalText() error = %v", err)
			}
			if got, want := out, tc.want; !cmp.Equal(got, want) {
				t.Fatalf("UnmarshalText() = %v, want %v", got, tc.want)
			}
		})
	}
}
