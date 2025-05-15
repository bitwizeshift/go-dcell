/*
Package errs is an internal package that provides shared errors across the
go-dcell library.

This package largely exists so that errors can be reused across both the
[expr], [dcell], and [funcs] packages without causing import cycles.
*/
package errs

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"strings"

	"rodusek.dev/pkg/dcell/internal/levenshtein"
)

var (
	// ErrEval is returned when an error occurs during evaluation of an
	// expression.
	ErrEval = errors.New("dcell eval")

	// ErrIncompatible is returned when two operands are incompatible.
	ErrIncompatible = errors.New("incompatible operands")

	// ErrUnknownName is an error returned when a name is not found in the
	// current context.
	ErrUnknownName = errors.New("unknown name")
)

// NameError is an error that indicates that a name does not exist in the
// current context. It provides suggestions for likely candidate names based on
// what is closest via Levenshtein distance.
type NameError struct {
	Input       string
	Suggestions []string
}

func (e *NameError) Error() string {
	var sb strings.Builder
	_, _ = fmt.Fprintf(&sb, "%v: '%s' does not exist", ErrUnknownName, e.Input)
	if len(e.Suggestions) == 1 {
		_, _ = fmt.Fprintf(&sb, ", did you mean '%s'?", e.Suggestions[0])
	}
	if len(e.Suggestions) > 1 {
		suggestions := strings.Join(e.Suggestions, "', '")
		_, _ = fmt.Fprintf(&sb, ", did you mean one of '%s'?", suggestions)
	}
	return sb.String()
}

func (e *NameError) Unwrap() error {
	return ErrUnknownName
}

// NewNameError creates a new [NameError] with the given input and options.
func NewNameError(input string, options iter.Seq[string]) *NameError {
	suggestions := levenshtein.Suggestions(input, options)
	slices.Sort(suggestions)
	suggestions = slices.Compact(suggestions)
	return &NameError{
		Input:       input,
		Suggestions: suggestions,
	}
}
