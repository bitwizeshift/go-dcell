package levenshtein_test

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/levenshtein"
)

func TestDamerauDistance(t *testing.T) {
	t.Parallel()
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
			t.Parallel()
			distance := levenshtein.DamerauDistance(tc.a, tc.b)

			if got, want := distance, tc.want; !cmp.Equal(got, want) {
				t.Errorf("DamereauDistance(%q, %q) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestSuggestion(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		input      string
		candidates []string
		want       []string
	}{
		{
			name:       "empty input",
			input:      "",
			candidates: []string{"apple", "banana", "orange"},
			want:       []string{},
		}, {
			name:       "match with 1 edit",
			input:      "red",
			candidates: []string{"red1", "blue1", "red2"},
			want:       []string{"red1", "red2"},
		}, {
			name:       "match with 2 edits",
			input:      "flawless",
			candidates: []string{"lawless", "lawn", "flawed"},
			want:       []string{"lawless"},
		}, {
			name:       "match with 3 edits",
			input:      "interface",
			candidates: []string{"inheritance", "internship", "interpolate", "interstate"},
			want:       []string{"interstate"},
		}, {
			name:       "match with 4+ edits for long input",
			input:      "instrumentations",
			candidates: []string{"interpretation", "instrumentation", "initialization", "imperative"},
			want:       []string{"instrumentation"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			suggestions := levenshtein.Suggestions(tc.input, slices.Values(tc.candidates))

			if got, want := suggestions, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Suggestions(%q, %v) = %v; want %v", tc.input, tc.candidates, got, want)
			}
		})
	}
}
