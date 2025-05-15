/*
Package arity is a micro-package for providing the [Arity] type, which is used
to identify whether the number of arguments being provided to an invocation is
valid.
*/
package arity

import (
	"errors"
	"fmt"
)

// ErrBadArity is an error that is returned when the arity of a function
// invocation is incorrect.
var ErrBadArity = errors.New("arity")

// Arity is an interface that checks the arity of function invocations.
type Arity interface {
	Check(n int) error
	arity()
}

type arity func(n int) error

func (a arity) Check(n int) error {
	return a(n)
}
func (a arity) arity() {}

// Exactly returns an Arity that checks for an exact number of arguments.
func Exactly(n int) Arity {
	return arity(func(m int) error {
		if m != n {
			return fmt.Errorf("%w: expected %d arg, got %d", ErrBadArity, n, m)
		}
		return nil
	})
}

// AtLeast returns an Arity that checks for at least a number of arguments.
func AtLeast(n int) Arity {
	return arity(func(m int) error {
		if m < n {
			return fmt.Errorf("%w: expected at least %d arg, got %d", ErrBadArity, n, m)
		}
		return nil
	})
}

// AtMost returns an Arity that checks for at most a number of arguments.
func AtMost(n int) Arity {
	return arity(func(m int) error {
		if m > n {
			return fmt.Errorf("%w: expected at most %d arg, got %d", ErrBadArity, n, m)
		}
		return nil
	})
}

// ClosedRange returns an Arity that checks for a range of arguments.
func ClosedRange(min, max int) Arity {
	return arity(func(n int) error {
		if n < min {
			return fmt.Errorf("%w: expected at least %d arg, got %d", ErrBadArity, min, n)
		}
		if n > max {
			return fmt.Errorf("%w: expected at most %d arg, got %d", ErrBadArity, max, n)
		}
		return nil
	})
}

// OneOf returns an Arity that checks for a set of arguments.
func OneOf(arities ...int) Arity {
	return arity(func(n int) error {
		for _, a := range arities {
			if n == a {
				return nil
			}
		}
		return fmt.Errorf("%w: expected one of %v arg, got %d", ErrBadArity, arities, n)
	})
}

// Any returns an Arity that checks for any number of arguments.
func Any() Arity {
	return arity(func(n int) error {
		return nil
	})
}

// None returns an Arity that checks for no arguments.
func None() Arity {
	return arity(func(n int) error {
		if n != 0 {
			return fmt.Errorf("%w: expected 0 args, got %d", ErrBadArity, n)
		}
		return nil
	})
}
