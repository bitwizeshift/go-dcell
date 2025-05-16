package compile

import (
	antlr "github.com/antlr4-go/antlr/v4"
)

// ErrorListener is both an Syntax and Parse error listener, which is used to
// accumulate any errors encountered as part of the compilation process.
type ErrorListener struct {
	*antlr.DefaultErrorListener

	// Errors is the list of errors encountered within this listener
	Errors []error
}

// SyntaxError is called when a syntax error is encountered during parsing.
func (c *ErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CompileError{
		Message: msg,
		Line:    line,
		Column:  column,
	})
}
