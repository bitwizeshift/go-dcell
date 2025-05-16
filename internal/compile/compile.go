package compile

import (
	"errors"
	"io"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/invocation"
	"rodusek.dev/pkg/dcell/internal/parser"
)

// Config provides compilation configuration to the [NewTree] function.
type Config struct {
	FuncTable *invocation.Table
}

// NewTree converts a string dcell expression into the proper Expression
// tree.
func NewTree(str string, cfg *Config) (expr.Expr, error) {
	return NewTreeFromReader(strings.NewReader(str), cfg)
}

// NewTreeFromReader converts a dcell expression from an io.Reader into the
// proper Expression tree.
func NewTreeFromReader(r io.Reader, cfg *Config) (expr.Expr, error) {
	input := antlr.NewIoStream(r)

	lexerErrors := &ErrorListener{}
	lexer := parser.NewDCellLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrors := &ErrorListener{}
	parser := parser.NewDCellParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	program := parser.Program()

	var errs []error
	errs = append(errs, lexerErrors.Errors...)
	errs = append(errs, parserErrors.Errors...)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	visitor := &Visitor{
		FuncTable: cfg.FuncTable,
	}
	return visitor.VisitProgram(program)
}
