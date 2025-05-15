package reflectcmp_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		name        string
		left, right reflect.Value
		want        bool
	}{
		{
			name:  "Integers, same type, equal",
			left:  reflect.ValueOf(42),
			right: reflect.ValueOf(42),
			want:  true,
		}, {
			name:  "Integers, different kind, equal",
			left:  reflect.ValueOf(uint8(42)),
			right: reflect.ValueOf(int64(42)),
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := reflectcmp.Equal(tc.left, tc.right)

			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Equal(%v, %v) = %v, want %v", tc.left, tc.right, got, want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		name        string
		left, right reflect.Value
		want        int
	}{
		{
			name:  "Integers, both signed, equal",
			left:  reflect.ValueOf(42),
			right: reflect.ValueOf(42),
			want:  0,
		}, {
			name:  "Integers, both unsigned, equal",
			left:  reflect.ValueOf(uint(42)),
			right: reflect.ValueOf(uint(42)),
			want:  0,
		}, {
			name:  "Integers, different sign, equal",
			left:  reflect.ValueOf(uint8(42)),
			right: reflect.ValueOf(int64(42)),
			want:  0,
		}, {
			name:  "Integers, signed < unsigned",
			left:  reflect.ValueOf(int8(1)),
			right: reflect.ValueOf(uint8(2)),
			want:  -1,
		}, {
			name:  "Integers, unsigned > signed",
			left:  reflect.ValueOf(uint16(100)),
			right: reflect.ValueOf(int8(10)),
			want:  1,
		}, {
			name:  "Floats, equal",
			left:  reflect.ValueOf(3.14),
			right: reflect.ValueOf(3.14),
			want:  0,
		}, {
			name:  "Floats, lhs < rhs",
			left:  reflect.ValueOf(2.71),
			right: reflect.ValueOf(3.14),
			want:  -1,
		}, {
			name:  "Floats, lhs > rhs",
			left:  reflect.ValueOf(4.0),
			right: reflect.ValueOf(3.14),
			want:  1,
		}, {
			name:  "String, equal",
			left:  reflect.ValueOf("foo"),
			right: reflect.ValueOf("foo"),
			want:  0,
		}, {
			name:  "String, lhs < rhs",
			left:  reflect.ValueOf("bar"),
			right: reflect.ValueOf("foo"),
			want:  -1,
		}, {
			name:  "String, lhs > rhs",
			left:  reflect.ValueOf("zoo"),
			right: reflect.ValueOf("foo"),
			want:  1,
		}, {
			name:  "Bool, equal",
			left:  reflect.ValueOf(true),
			right: reflect.ValueOf(true),
			want:  0,
		}, {
			name:  "Bool, lhs < rhs",
			left:  reflect.ValueOf(false),
			right: reflect.ValueOf(true),
			want:  -1,
		}, {
			name:  "Bool, lhs > rhs",
			left:  reflect.ValueOf(true),
			right: reflect.ValueOf(false),
			want:  1,
		}, {
			name:  "Struct, equal",
			left:  reflect.ValueOf(struct{ A int }{A: 1}),
			right: reflect.ValueOf(struct{ A int }{A: 1}),
			want:  0,
		}, {
			name:  "Struct, lhs < rhs (field value)",
			left:  reflect.ValueOf(struct{ A int }{A: 1}),
			right: reflect.ValueOf(struct{ A int }{A: 2}),
			want:  -1,
		}, {
			name:  "Struct, lhs > rhs (field value)",
			left:  reflect.ValueOf(struct{ A int }{A: 3}),
			right: reflect.ValueOf(struct{ A int }{A: 2}),
			want:  1,
		}, {
			name:  "Struct, different types",
			left:  reflect.ValueOf(struct{ A int }{A: 1}),
			right: reflect.ValueOf(struct{ B int }{B: 1}),
			want:  strings.Compare(reflect.TypeOf(struct{ A int }{A: 1}).String(), reflect.TypeOf(struct{ B int }{B: 1}).String()),
		}, {
			name:  "Slice, equal",
			left:  reflect.ValueOf([]int{1, 2, 3}),
			right: reflect.ValueOf([]int{1, 2, 3}),
			want:  0,
		}, {
			name:  "Slice, lhs < rhs (length)",
			left:  reflect.ValueOf([]int{1, 2}),
			right: reflect.ValueOf([]int{1, 2, 3}),
			want:  -1,
		}, {
			name:  "Slice, lhs > rhs (element)",
			left:  reflect.ValueOf([]int{1, 3, 3}),
			right: reflect.ValueOf([]int{1, 2, 3}),
			want:  1,
		}, {
			name:  "Map, equal (single key)",
			left:  reflect.ValueOf(map[string]int{"a": 1}),
			right: reflect.ValueOf(map[string]int{"a": 1}),
			want:  0,
		}, {
			name:  "Map, lhs < rhs (length)",
			left:  reflect.ValueOf(map[string]int{"a": 1}),
			right: reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			want:  -1,
		}, {
			name:  "Map, lhs > rhs (length)",
			left:  reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			right: reflect.ValueOf(map[string]int{"a": 1}),
			want:  1,
		}, {
			name:  "Map, different values (single key)",
			left:  reflect.ValueOf(map[string]int{"a": 2}),
			right: reflect.ValueOf(map[string]int{"a": 1}),
			want:  1,
		}, {
			name:  "Map, different keys",
			left:  reflect.ValueOf(map[string]int{"a": 1}),
			right: reflect.ValueOf(map[string]int{"b": 1}),
			want:  strings.Compare("a", "b"),
		}, {
			name:  "Nil vs Nil",
			left:  reflect.Value{},
			right: reflect.Value{},
			want:  0,
		}, {
			name:  "Nil vs Int",
			left:  reflect.Value{},
			right: reflect.ValueOf(1),
			want:  1,
		}, {
			name:  "Int vs Nil",
			left:  reflect.ValueOf(1),
			right: reflect.Value{},
			want:  -1,
		}, {
			name:  "Unknown type (channels)",
			left:  reflect.ValueOf(make(chan int)),
			right: reflect.ValueOf(make(chan int)),
			want:  1, // rtypeUnknown always returns 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := reflectcmp.Compare(tc.left, tc.right)

			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Compare(%v, %v) = %v, want %v", tc.left, tc.right, got, want)
			}
		})
	}
}
