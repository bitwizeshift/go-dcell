package levenshtein_test

import (
	"testing"

	"rodusek.dev/pkg/dcell/internal/levenshtein"
)

func TestDamerauDistance(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "a is empty",
			a:    "",
			b:    "abc",
			want: 3,
		}, {
			name: "b is empty",
			a:    "abc",
			b:    "",
			want: 3,
		}, {
			name: "transposition",
			a:    "abcde",
			b:    "acbde",
			want: 1,
		}, {
			name: "insertion",
			a:    "abc",
			b:    "abdc",
			want: 1,
		}, {
			name: "deletion",
			a:    "abc",
			b:    "ab",
			want: 1,
		}, {
			name: "substitution",
			a:    "abc",
			b:    "abd",
			want: 1,
		}, {
			name: "different lengths",
			a:    "kitten",
			b:    "sitting",
			want: 3,
		}, {
			name: "one character difference",
			a:    "flaw",
			b:    "lawn",
			want: 2,
		}, {
			name: "multiple edits",
			a:    "intention",
			b:    "execution",
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distance := levenshtein.DamerauDistance(tc.a, tc.b)

			if got, want := distance, tc.want; got != want {
				t.Errorf("DamereauDistance(%q, %q) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
