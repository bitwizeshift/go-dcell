package intcmp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
	"rodusek.dev/pkg/dcell/internal/intcmp"
)

func compare[L, R constraints.Integer](l L, r R) func() int {
	return func() int {
		return intcmp.Compare(l, r)
	}
}

func equal[L, R constraints.Integer](l L, r R) func() bool {
	return func() bool {
		return intcmp.Equal(l, r)
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		test func() int
		want int
	}{
		{
			name: "Lhs and Rhs are signed, equal",
			test: compare(1, 1),
			want: 0,
		}, {
			name: "Lhs and Rhs are signed, less",
			test: compare(1, 2),
			want: -1,
		}, {
			name: "Lhs and Rhs are signed, greater",
			test: compare(2, 1),
			want: 1,
		}, {
			name: "Lhs and Rhs are unsigned, equal",
			test: compare(uint8(1), uint8(1)),
			want: 0,
		}, {
			name: "Lhs and Rhs are unsigned, less",
			test: compare(uint8(1), uint8(2)),
			want: -1,
		}, {
			name: "Lhs and Rhs are unsigned, greater",
			test: compare(uint8(2), uint8(1)),
			want: 1,
		}, {
			name: "Lhs is signed and Rhs is unsigned, equal",
			test: compare(int8(1), uint8(1)),
			want: 0,
		}, {
			name: "Lhs is signed and Rhs is unsigned, less",
			test: compare(-1, uint8(2)),
			want: -1,
		}, {
			name: "Lhs is signed and Rhs is unsigned, greater",
			test: compare(int8(2), uint8(1)),
			want: 1,
		}, {
			name: "Lhs is unsigned and Rhs is signed, equal",
			test: compare(uint8(1), int8(1)),
			want: 0,
		}, {
			name: "Lhs is unsigned and Rhs is signed, less",
			test: compare(uint8(1), int8(2)),
			want: -1,
		}, {
			name: "Lhs is unsigned and Rhs is signed, greater",
			test: compare(uint8(2), -1),
			want: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := tc.test()

			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Compare(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		test func() bool
		want bool
	}{
		{
			name: "Lhs and Rhs are signed, equal",
			test: equal(1, 1),
			want: true,
		}, {
			name: "Lhs and Rhs are signed, not equal",
			test: equal(1, 2),
			want: false,
		}, {
			name: "Lhs and Rhs are unsigned, equal",
			test: equal(uint8(1), uint8(1)),
			want: true,
		}, {
			name: "Lhs and Rhs are unsigned, not equal",
			test: equal(uint8(1), uint8(2)),
			want: false,
		}, {
			name: "Lhs is signed and Rhs is unsigned, equal",
			test: equal(int8(1), uint8(1)),
			want: true,
		}, {
			name: "Lhs is signed and Rhs is unsigned, not equal",
			test: equal(-1, uint8(2)),
			want: false,
		}, {
			name: "Lhs is unsigned and Rhs is signed, equal",
			test: equal(uint8(1), int8(1)),
			want: true,
		}, {
			name: "Lhs is unsigned and Rhs is signed, not equal",
			test: equal(uint8(1), int8(2)),
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := tc.test()

			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Equal(...) = %v, want %v", got, want)
			}
		})
	}
}
