package exprtest_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestBoolean(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input bool
	}{
		{
			name:  "true",
			input: true,
		},
		{
			name:  "false",
			input: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := exprtest.Boolean(tc.input)

			result, err := sut.Eval(nil)

			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Boolean() error = %v, want %v", got, want)
			}
			if got, want := result, reflect.ValueOf(tc.input); !reflectcmp.Equal(got, want) {
				t.Fatalf("Boolean() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestInteger(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input int
	}{
		{
			name:  "zero",
			input: 0,
		},
		{
			name:  "positive",
			input: 42,
		},
		{
			name:  "negative",
			input: -42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := exprtest.Integer(tc.input)

			result, err := sut.Eval(nil)

			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Integer() error = %v, want %v", got, want)
			}
			if got, want := result, reflect.ValueOf(tc.input); !reflectcmp.Equal(got, want) {
				t.Fatalf("Integer() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input float64
	}{
		{
			name:  "zero",
			input: 0.0,
		},
		{
			name:  "positive",
			input: 42.0,
		},
		{
			name:  "negative",
			input: -42.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := exprtest.Float(tc.input)

			result, err := sut.Eval(nil)

			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Float() error = %v, want %v", got, want)
			}
			if got, want := result, reflect.ValueOf(tc.input); !reflectcmp.Equal(got, want) {
				t.Fatalf("Float() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "empty",
			input: "",
		},
		{
			name:  "hello",
			input: "hello",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := exprtest.String(tc.input)

			result, err := sut.Eval(nil)

			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("String() error = %v, want %v", got, want)
			}
			if got, want := result, reflect.ValueOf(tc.input); !reflectcmp.Equal(got, want) {
				t.Fatalf("String() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty",
			input: []int{},
		},
		{
			name:  "one element",
			input: []int{1},
		},
		{
			name:  "multiple elements",
			input: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := exprtest.Slice(tc.input...)

			result, err := sut.Eval(nil)

			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Slice() error = %v, want %v", got, want)
			}
			if got, want := result, reflect.ValueOf(tc.input); !reflectcmp.Equal(got, want) {
				t.Fatalf("Slice() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestError(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	sut := exprtest.Error(testErr)

	result, err := sut.Eval(nil)

	if got, want := err, testErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Error() error = %v, want %v", got, want)
	}
	if got, want := result, (reflect.Value{}); !reflectcmp.Equal(got, want) {
		t.Errorf("Error() = %v, want %v", got.Interface(), want.Interface())
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	sut := exprtest.Empty()

	result, err := sut.Eval(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Empty() error = %v, want %v", got, want)
	}
	if got, want := result, (reflect.Value{}); !reflectcmp.Equal(got, want) {
		t.Errorf("Empty() = %v, want %v", got.Interface(), want.Interface())
	}
}
