package expr_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
)

func TestWildcardExpr(t *testing.T) {
	t.Parallel()

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
			name:  "Input is pointer to nil pointer",
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
			name: "Input is struct with fields of the same type",
			input: struct {
				FieldOne string `dcell:"field-one"`
				FieldTwo string `dcell:"field-two"`
			}{
				FieldOne: "value1",
				FieldTwo: "value2",
			},
			want: []string{"value1", "value2"},
		}, {
			name: "Input is struct with fields of different types",
			input: struct {
				unexported int
				FieldOne   string
				FieldTwo   int
			}{
				FieldOne: "value1",
				FieldTwo: 42,
			},
			want: []any{"value1", 42},
		}, {
			name:  "Input is empty struct",
			input: struct{}{},
			want:  nil,
		}, {
			name: "Input is map with values of different types",
			input: map[string]any{
				"key1": "value1",
				"key2": 42,
			},
			want: []any{"value1", 42},
		}, {
			name:    "Input is not struct or map",
			input:   123,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Wildcard()
			input := expr.NewContext(reflect.ValueOf(tc.input))
			var expect reflect.Value
			if tc.want != nil {
				expect = reflect.ValueOf(tc.want)
			}

			got, err := sut.Eval(input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Wildcard.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, expect; !reflectEqual(got, want) {
				t.Errorf("Wildcard.Eval() = %v, want %v", got, want)
			}
		})
	}
}

func reflectEqual(got, want reflect.Value) bool {
	if got.Kind() == reflect.Slice && want.Kind() == reflect.Slice {
		return cmp.Equal(got.Interface(), want.Interface())
	}
	return cmp.Equal(got, want)
}
