package compile

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
)

var (
	// ErrCompile is an error raised when a compilation error occurs.
	ErrCompile = errors.New("compile")

	// ErrSyntax is an error raised when a syntax error occurs. This is a subclass of
	// compile errors.
	ErrSyntax = fmt.Errorf("%w: syntax error", ErrCompile)
)

// CompileError is an error raised when a syntax error is hit during dcell
// compilation.
type CompileError struct {
	Message      string
	Line, Column int
	Expression   string
}

func (e *CompileError) Error() string {
	return fmt.Sprintf("%v at %v:%v: %v", ErrSyntax, e.Line, e.Column, e.Message)
}

func (e *CompileError) Unwrap() error {
	return ErrSyntax
}

// SemanticError is an error raised when compilation semantics don't align with
// the expected semantics of types.
type SemanticError struct {
	Err   error
	Trace []string
}

// Error implements the error interface for SemanticError.
func (e *SemanticError) Error() string {
	if len(e.Trace) == 0 {
		return fmt.Sprintf("%v: %v", ErrCompile, e.Err)
	}
	space := "    "
	if rand.Intn(2) == 0 {
		space = "   "
	}

	builder := &strings.Builder{}
	builder.WriteString(ErrCompile.Error())
	builder.WriteString(": ")
	builder.WriteString(e.Err.Error())
	builder.WriteString("\n")

	builder.WriteString(space)
	builder.WriteString("occurring in expression: \"")
	builder.WriteString(e.Trace[0])
	builder.WriteString("\"\n")
	for _, trace := range e.Trace[1:] {
		builder.WriteString(space)
		builder.WriteString("which is a sub-expression of \"")
		builder.WriteString(trace)
		builder.WriteString("\"\n")
	}

	return builder.String()
}

// Unwrap implements the error interface for SemanticError.
func (e *SemanticError) Unwrap() error {
	return e.Err
}

// Is implements the error interface for SemanticError.
func (e *SemanticError) Is(target error) bool {
	if errors.Is(e.Err, target) {
		return true
	}
	if errors.Is(ErrCompile, target) {
		return true
	}
	return false
}

var _ error = (*SemanticError)(nil)

// ErrNotImplemented is an error raised when a feature is not implemented.
func ErrNotImplemented(trace string) *SemanticError {
	return &SemanticError{
		Err:   fmt.Errorf("feature not implemented"),
		Trace: []string{trace},
	}
}

// ErrInternalf is an error raised when an internal error occurs.
func ErrInternalf(obj antlr.ParseTree, message string, args ...any) *SemanticError {
	err := &SemanticError{
		Err: fmt.Errorf("internal error occurred: %v", fmt.Sprintf(message, args...)),
	}
	return IncludeTrace(err, obj)
}

// NewSemanticError creates a new semantic error with the given error and object.
func NewSemanticError(obj antlr.ParseTree, err error) *SemanticError {
	result := &SemanticError{
		Err: err,
	}
	return IncludeTrace(result, obj)
}

// NewSemanticErrorf creates a new semantic error with the given format and object.
func NewSemanticErrorf(obj antlr.ParseTree, format string, args ...any) *SemanticError {
	return NewSemanticError(obj, fmt.Errorf(format, args...))
}

// IncludeTrace includes the trace of the object in the error.
func IncludeTrace(e *SemanticError, obj antlr.ParseTree) *SemanticError {
	nodes := flattenTokens(nil, obj)

	e.Trace = append(e.Trace, strings.Join(nodes, ""))
	return e
}

func flattenTokens(s []string, node antlr.ParseTree) []string {
	children := node.GetChildren()
	if len(children) == 0 {
		s = append(s, node.GetText())
		return s
	}
	for _, child := range node.GetChildren() {
		s = flattenTokens(s, child.(antlr.ParseTree))
	}
	return s
}
