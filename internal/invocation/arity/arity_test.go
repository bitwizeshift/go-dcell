package arity_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/invocation/arity"
)

func TestArity(t *testing.T) {
	testCases := []struct {
		name  string
		arity arity.Arity
		args  int
		want  error
	}{
		{
			name:  "Exactly: correct arity",
			arity: arity.Exactly(2),
			args:  2,
		}, {
			name:  "Exactly: incorrect arity",
			arity: arity.Exactly(2),
			args:  3,
			want:  arity.ErrBadArity,
		}, {
			name:  "AtLeast: minimum allowed",
			arity: arity.AtLeast(2),
			args:  2,
		}, {
			name:  "AtLeast: more than minimum",
			arity: arity.AtLeast(2),
			args:  3,
		}, {
			name:  "AtLeast: less than minimum",
			arity: arity.AtLeast(2),
			args:  1,
			want:  arity.ErrBadArity,
		}, {
			name:  "AtMost: maximum allowed",
			arity: arity.AtMost(2),
			args:  2,
		}, {
			name:  "AtMost: less than maximum",
			arity: arity.AtMost(2),
			args:  1,
		}, {
			name:  "AtMost: more than maximum",
			arity: arity.AtMost(2),
			args:  3,
			want:  arity.ErrBadArity,
		}, {
			name:  "ClosedRange: within range",
			arity: arity.ClosedRange(2, 4),
			args:  3,
		}, {
			name:  "ClosedRange: minimum",
			arity: arity.ClosedRange(2, 4),
			args:  2,
		}, {
			name:  "ClosedRange: maximum",
			arity: arity.ClosedRange(2, 4),
			args:  4,
		}, {
			name:  "ClosedRange: less than minimum",
			arity: arity.ClosedRange(2, 4),
			args:  1,
			want:  arity.ErrBadArity,
		}, {
			name:  "ClosedRange: more than maximum",
			arity: arity.ClosedRange(2, 4),
			args:  5,
			want:  arity.ErrBadArity,
		}, {
			name:  "OneOf: valid arity",
			arity: arity.OneOf(2, 4, 6),
			args:  4,
		}, {
			name:  "OneOf: invalid arity",
			arity: arity.OneOf(2, 4, 6),
			args:  3,
			want:  arity.ErrBadArity,
		}, {
			name:  "Any: any arity",
			arity: arity.Any(),
			args:  0,
		}, {
			name:  "None: no arity",
			arity: arity.None(),
			args:  0,
		}, {
			name:  "None: some arity",
			arity: arity.None(),
			args:  1,
			want:  arity.ErrBadArity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.arity.Check(tc.args)

			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("CheckArity() error = %v, want %v", tc.want, got)
			}
		})
	}
}
