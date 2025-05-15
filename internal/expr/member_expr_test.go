package expr_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func ptr[T any](v T) *T {
	return &v
}

func TestMemberExpr(t *testing.T) {
	t.Parallel()

	const member = "foo"
	testCases := []struct {
		name    string
		input   any
		want    any
		wantErr error
	}{
		{
			name:  "Input is nil pointer",
			input: (*string)(nil),
			want:  nil,
		}, {
			name:  "input is pointer to nil pointer",
			input: ptr((*string)(nil)),
			want:  nil,
		}, {
			name:  "Input is nil map",
			input: map[string]any(nil),
			want:  nil,
		}, {
			name:  "Input is pointer to nil map",
			input: ptr(map[string]any(nil)),
			want:  nil,
		}, {
			name:  "Input is map, field exists",
			input: map[string]any{"foo": "bar"},
			want:  "bar",
		}, {
			name:    "Input is map, field does not exist",
			input:   map[string]any{"bar": "baz"},
			wantErr: errs.ErrUnknownName,
		}, {
			name:    "Input is map, key is not a string",
			input:   map[any]any{1: "bar"},
			wantErr: errs.ErrUnknownName,
		}, {
			name: "Input is struct, field exists",
			input: struct {
				somethingUnexported any
				NoTag               string
				WrongKey            string `dcell:"wrong-key"`
				Foo                 string `dcell:"foo"`
			}{
				Foo: "bar",
			},
			want: "bar",
		}, {
			name:  "Input is empty struct",
			input: struct{}{},
			want:  nil,
		}, {
			name: "Input is struct, field does not exist",
			input: struct {
				KeyOne string `dcell:"key-one"`
				KeyTwo string `dcell:"key-two"`
				Foop   string `dcell:"foop"`
				Foob   string `dcell:"foob"`
			}{},
			wantErr: errs.ErrUnknownName,
		}, {
			name:    "Input is not struct or map",
			input:   123,
			wantErr: errs.ErrUnknownName,
		}, {
			name: "Input is slice of struct, field exists",
			input: []*struct {
				somethingUnexported any
				Foo                 string `dcell:"foo"`
			}{
				{
					Foo: "bar",
				},
			},
			want: []string{"bar"},
		}, {
			name: "Input is slice of map, field exists",
			input: []map[string]any{
				{
					"foo": "bar",
				},
			},
			want: []string{"bar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := expr.Member(member)
			input := expr.NewContext(reflect.ValueOf(tc.input))
			var expect reflect.Value
			if tc.want != nil {
				expect = reflect.ValueOf(tc.want)
			}

			got, err := sut.Eval(input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("MemberEval(%q) error = %v, want %v", member, got, want)
			}
			if got, want := got, expect; !reflectcmp.Equal(got, want) {
				t.Errorf("MemberEval(%q) = %v, want %v", member, got, want)
			}
		})
	}
}
