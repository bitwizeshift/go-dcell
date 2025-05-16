package compile

import (
	"strconv"

	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/invocation"
	"rodusek.dev/pkg/dcell/internal/parser"
)

// Visitor is a visitor that walks the parse tree to generate an expression
type Visitor struct {
	FuncTable *invocation.Table
}

// VisitProgram visits the root of the parse tree
func (v *Visitor) VisitProgram(ctx parser.IProgramContext) (expr.Expr, error) {
	return v.visitExpression(ctx.Expression())
}

//------------------------------------------------------------------------------
// Expressions
//------------------------------------------------------------------------------

func (v *Visitor) visitExpression(ctx parser.IExpressionContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.TermExpressionContext:
		return v.visitTermExpression(ctx)
	case *parser.InvocationExpressionContext:
		return v.visitInvocationExpression(ctx)
	case *parser.IndexExpressionContext:
		return v.visitIndexExpression(ctx)
	case *parser.ParenthesisExpressionContext:
		return v.visitParenthesisExpression(ctx)
	case *parser.LogicalNotExpressionContext:
		return v.visitLogicalNotExpression(ctx)
	case *parser.BitwiseNotExpressionContext:
		return v.visitBitwiseNotExpression(ctx)
	case *parser.PolarityExpressionContext:
		return v.visitPolarityExpression(ctx)
	case *parser.ExponentiationExpressionContext:
		return v.visitExponentiationExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return v.visitMultiplicativeExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return v.visitAdditiveExpression(ctx)
	case *parser.LogicalAndExpressionContext:
		return v.visitLogicalAndExpression(ctx)
	case *parser.LogicalOrExpressionContext:
		return v.visitLogicalOrExpression(ctx)
	case *parser.ImplicationExpressionContext:
		return v.visitImplicationExpression(ctx)
	case *parser.ShiftExpressionContext:
		return v.visitShiftExpression(ctx)
	case *parser.BitwiseAndExpressionContext:
		return v.visitBitwiseAndExpression(ctx)
	case *parser.BitwiseOrExpressionContext:
		return v.visitBitwiseOrExpression(ctx)
	case *parser.InequalityExpressionContext:
		return v.visitInequalityExpression(ctx)
	case *parser.EqualityExpressionContext:
		return v.visitEqualityExpression(ctx)
	case *parser.TernaryExpressionContext:
		return v.visitTernaryExpression(ctx)
	case *parser.ElvisExpressionContext:
		return v.visitElvisExpression(ctx)
	case *parser.CoalesceExpressionContext:
		return v.visitCoalesceExpression(ctx)
	case *parser.IsExpressionContext:
		return v.visitIsExpression(ctx)
	case *parser.CastExpressionContext:
		return v.visitCastExpression(ctx)
	case *parser.ContainsExpressionContext:
		return v.visitContainsExpression(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected expression type: %T", ctx)
}

func (v *Visitor) visitTermExpression(ctx *parser.TermExpressionContext) (expr.Expr, error) {
	return v.visitTerm(ctx.Term())
}

func (v *Visitor) visitInvocationExpression(ctx *parser.InvocationExpressionContext) (expr.Expr, error) {
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	right, err := v.visitInvocation(ctx.Invocation(), false)
	if err != nil {
		return nil, err
	}
	return expr.Sequence(left, right), nil
}

func (v *Visitor) visitIndexExpression(ctx *parser.IndexExpressionContext) (expr.Expr, error) {
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	right, err := v.visitIndex(ctx.Index())
	return expr.Sequence(left, right), nil
}

func (v *Visitor) visitParenthesisExpression(ctx *parser.ParenthesisExpressionContext) (expr.Expr, error) {
	return v.visitExpression(ctx.Expression())
}

func (v *Visitor) visitLogicalNotExpression(ctx *parser.LogicalNotExpressionContext) (expr.Expr, error) {
	rest, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return expr.LogicalNot(rest), nil
}

func (v *Visitor) visitBitwiseNotExpression(ctx *parser.BitwiseNotExpressionContext) (expr.Expr, error) {
	rest, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return expr.BitwiseNot(rest), nil
}

func (v *Visitor) visitPolarityExpression(ctx *parser.PolarityExpressionContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(0))
	switch op {
	case "+":
		return expr.PolarityPlus(e), nil
	case "-":
		return expr.PolarityMinus(e), nil
	}
	return nil, ErrInternalf(ctx, "unexpected polarity operator: %s", op)
}

func (v *Visitor) visitExponentiationExpression(ctx *parser.ExponentiationExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	return expr.Power(left, right), nil
}

func (v *Visitor) visitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]

	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "*":
		return expr.Multiply(left, right), nil
	case "/":
		return expr.Divide(left, right), nil
	case "//":
		return expr.FloorDivide(left, right), nil
	case "%":
		return expr.Modulus(left, right), nil
	}
	return nil, ErrInternalf(ctx, "unexpected multiplicative operator: %s", op)
}

func (v *Visitor) visitAdditiveExpression(ctx *parser.AdditiveExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]

	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "+":
		return expr.Add(left, right), nil
	case "-":
		return expr.Subtract(left, right), nil
	}
	return nil, ErrInternalf(ctx, "additive expression %q not implemented", op)
}

func (v *Visitor) visitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]

	return expr.LogicalAnd(left, right), nil
}

func (v *Visitor) visitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]

	return expr.LogicalOr(left, right), nil
}

func (v *Visitor) visitImplicationExpression(ctx *parser.ImplicationExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	return expr.Implies(left, right), nil
}

func (v *Visitor) visitShiftExpression(ctx *parser.ShiftExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "<<":
		return expr.BitwiseShiftLeft(left, right), nil
	case ">>":
		return expr.BitwiseShiftRight(left, right), nil
	}
	return nil, ErrInternalf(ctx, "bitwise shift expression %q not implemented", op)
}

func (v *Visitor) visitBitwiseAndExpression(ctx *parser.BitwiseAndExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	return expr.BitwiseAnd(left, right), nil
}

func (v *Visitor) visitBitwiseOrExpression(ctx *parser.BitwiseOrExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "|":
		return expr.BitwiseOr(left, right), nil
	case "^":
		return expr.BitwiseXor(left, right), nil
	}
	return nil, ErrInternalf(ctx, "bitwise expression %q not implemented", op)
}

func (v *Visitor) visitInequalityExpression(ctx *parser.InequalityExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "<":
		return expr.LessThan(left, right), nil
	case "<=":
		return expr.LessThanOrEqual(left, right), nil
	case ">":
		return expr.GreaterThan(left, right), nil
	case ">=":
		return expr.GreaterThanOrEqual(left, right), nil
	}
	return nil, ErrInternalf(ctx, "inequality expression %q not implemented", op)
}

func (v *Visitor) visitEqualityExpression(ctx *parser.EqualityExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "==":
		return expr.Equal(left, right), nil
	case "!=":
		return expr.NotEqual(left, right), nil
	}
	return nil, ErrInternalf(ctx, "equality expression %q not implemented", op)
}

func (v *Visitor) visitTernaryExpression(ctx *parser.TernaryExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	condition, trueExpr, falseExpr := exprs[0], exprs[1], exprs[2]
	return expr.Ternary(condition, trueExpr, falseExpr), nil
}

func (v *Visitor) visitElvisExpression(ctx *parser.ElvisExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	condition, falseExpr := exprs[0], exprs[1]
	return expr.Elvis(condition, falseExpr), nil
}

func (v *Visitor) visitCoalesceExpression(ctx *parser.CoalesceExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	return expr.Coalesce(left, right), nil
}

func (v *Visitor) visitIsExpression(ctx *parser.IsExpressionContext) (expr.Expr, error) {
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	right, err := v.visitType(ctx.Type_())
	if err != nil {
		return nil, err
	}
	condition := v.getTreeText(ctx.GetChild(2))
	if condition == "not" {
		return expr.IsNot(left, right), nil
	}
	return expr.Is(left, right), nil
}

func (v *Visitor) visitCastExpression(ctx *parser.CastExpressionContext) (expr.Expr, error) {
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	right, err := v.visitType(ctx.Type_())
	if err != nil {
		return nil, err
	}
	return expr.As(left, right), nil
}

func (v *Visitor) visitContainsExpression(ctx *parser.ContainsExpressionContext) (expr.Expr, error) {
	exprs, err := v.visitExpressions(ctx.AllExpression())
	if err != nil {
		return nil, err
	}
	left, right := exprs[0], exprs[1]
	condition := v.getTreeText(ctx.GetChild(1))
	if condition == "in" {
		return expr.In(left, right), nil
	}
	return expr.NotIn(left, right), nil
}

func (v *Visitor) visitExpressions(ctxs []parser.IExpressionContext) ([]expr.Expr, error) {
	var exprs []expr.Expr
	for _, ctx := range ctxs {
		expr, err := v.visitExpression(ctx)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
	}
	return exprs, nil
}

//------------------------------------------------------------------------------
// Terms
//------------------------------------------------------------------------------

func (v *Visitor) visitTerm(ctx parser.ITermContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.LiteralTermContext:
		return v.visitLiteralTerm(ctx)
	case *parser.InvocationTermContext:
		return v.visitInvocationTerm(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected term type: %T", ctx)
}

func (v *Visitor) visitLiteralTerm(ctx *parser.LiteralTermContext) (expr.Expr, error) {
	literal, err := v.visitLiteral(ctx.Literal())
	if err != nil {
		return nil, err
	}
	return expr.Literal(literal), nil
}

func (v *Visitor) visitInvocationTerm(ctx *parser.InvocationTermContext) (expr.Expr, error) {
	return v.visitInvocation(ctx.Invocation(), true)
}

//------------------------------------------------------------------------------
// Invocations
//------------------------------------------------------------------------------

func (v *Visitor) visitInvocation(ctx parser.IInvocationContext, isRoot bool) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.FunctionInvocationContext:
		return v.visitFunctionInvocation(ctx, isRoot)
	case *parser.WildcardInvocationContext:
		return v.visitWildcardInvocation(ctx)
	case *parser.MemberInvocationContext:
		return v.visitMemberInvocation(ctx), nil
	}
	return nil, ErrInternalf(ctx, "unexpected invocation type: %T", ctx)
}

func (v *Visitor) visitFunctionInvocation(ctx *parser.FunctionInvocationContext, isRoot bool) (expr.Expr, error) {
	funcName := v.visitIdentifier(ctx.Identifier())

	params, err := v.visitParameterList(ctx.ParameterList())
	if err != nil {
		return nil, err
	}
	entry, ok := v.FuncTable.Lookup(funcName)
	if !ok {
		err := errs.NewNameError(funcName, v.FuncTable.FunctionNames())
		return nil, NewSemanticErrorf(ctx, "%w", err)
	}

	args := len(params)
	if !isRoot {
		args++ // member funcs include the root as the first arg
	}
	if err := entry.TestArity(args); err != nil {
		return nil, err
	}

	if isRoot {
		return expr.FreeFunc(entry.Invoke, params...), nil
	}
	return expr.MemberFunc(entry.Invoke, params...), nil
}

func (v *Visitor) visitWildcardInvocation(*parser.WildcardInvocationContext) (expr.Expr, error) {
	return expr.Wildcard(), nil
}

func (v *Visitor) visitMemberInvocation(ctx *parser.MemberInvocationContext) expr.Expr {
	memberName := v.visitIdentifier(ctx.Identifier())

	return expr.MemberExpr(memberName)
}

func (v *Visitor) visitParameterList(ctx parser.IParameterListContext) ([]expr.Expr, error) {
	if ctx == nil {
		return nil, nil
	}
	var params []expr.Expr
	for _, param := range ctx.AllExpression() {
		expr, err := v.visitExpression(param)
		if err != nil {
			return nil, err
		}
		params = append(params, expr)
	}
	return params, nil
}

func (v *Visitor) visitIdentifier(ctx parser.IIdentifierContext) string {
	return ctx.GetText()
}

//------------------------------------------------------------------------------
// Index
//------------------------------------------------------------------------------

func (v *Visitor) visitIndex(ctx parser.IIndexContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.SliceIndexContext:
		return v.visitSliceIndex(ctx)
	case *parser.ExpressionIndexContext:
		return v.visitExpressionIndex(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected index param type: %T", ctx)
}

func (v *Visitor) visitSliceIndex(ctx *parser.SliceIndexContext) (expr.Expr, error) {
	var low, high expr.Expr
	low = expr.Literal(0)
	if e := ctx.Expression(0); e != nil {
		expr, err := v.visitExpression(e)
		if err != nil {
			return nil, err
		}
		low = expr
	}
	if e := ctx.Expression(1); e != nil {
		expr, err := v.visitExpression(e)
		if err != nil {
			return nil, err
		}
		high = expr
	}

	return expr.IndexSlice(low, high), nil
}

func (v *Visitor) visitExpressionIndex(ctx *parser.ExpressionIndexContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return expr.Index(e), nil
}

//------------------------------------------------------------------------------
// Literal
//------------------------------------------------------------------------------

func (v *Visitor) visitLiteral(ctx parser.ILiteralContext) (any, error) {
	switch ctx := ctx.(type) {
	case *parser.StringLiteralContext:
		return v.visitStringLiteral(ctx)
	case *parser.IntegerLiteralContext:
		return v.visitIntegerLiteral(ctx)
	case *parser.FloatLiteralContext:
		return v.visitFloatLiteral(ctx)
	case *parser.BooleanLiteralContext:
		return v.visitBooleanLiteral(ctx), nil
	case *parser.NullLiteralContext:
		return v.visitNullLiteral(ctx), nil
	case *parser.ListLiteralContext:
		return v.visitListLiteral(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected literal type: %T", ctx)
}

func (v *Visitor) visitStringLiteral(ctx *parser.StringLiteralContext) (string, error) {
	var str string
	switch ctx := ctx.String_().(type) {
	case *parser.SingleQuoteStringContext:
		raw := ctx.SINGLE_QUOTE_STRING().GetText()
		str = "\"" + raw[1:len(raw)-1] + "\""
	case *parser.DoubleQuoteStringContext:
		str = ctx.DOUBLE_QUOTE_STRING().GetText()
	case *parser.TripleQuoteStringContext:
		raw := ctx.TRIPLE_QUOTE_STRING().GetText()
		str = "\"" + raw[3:len(raw)-3] + "\""
	}
	return strconv.Unquote(str)
}

func (v *Visitor) visitIntegerLiteral(ctx *parser.IntegerLiteralContext) (int64, error) {
	var str string
	var base int
	switch ctx := ctx.Integer().(type) {
	case *parser.DecimalIntegerContext:
		str = ctx.DECIMAL_INTEGER().GetText()
		base = 10
	case *parser.HexIntegerContext:
		raw := ctx.HEX_INTEGER().GetText()
		str = raw[2:] // remove 0x
		base = 16
	case *parser.OctalIntegerContext:
		raw := ctx.OCTAL_INTEGER().GetText()
		str = raw[2:] // remove 0o
		base = 8
	case *parser.BinaryIntegerContext:
		raw := ctx.BINARY_INTEGER().GetText()
		str = raw[2:] // remove 0b
		base = 2
	}
	return strconv.ParseInt(str, base, 64)
}

func (v *Visitor) visitFloatLiteral(ctx *parser.FloatLiteralContext) (float64, error) {
	var str string
	switch ctx := ctx.Float().(type) {
	case *parser.DecimalFloatContext:
		str = ctx.DECIMAL_FLOAT().GetText()
	case *parser.ScientificFloatContext:
		str = ctx.SCIENTIFIC_FLOAT().GetText()
	}
	return strconv.ParseFloat(str, 64)
}

func (v *Visitor) visitBooleanLiteral(ctx *parser.BooleanLiteralContext) bool {
	return ctx.GetText() == "true"
}

func (v *Visitor) visitNullLiteral(ctx *parser.NullLiteralContext) any {
	_ = ctx
	return nil
}

func (v *Visitor) visitListLiteral(ctx *parser.ListLiteralContext) ([]any, error) {
	var result []any
	for _, item := range ctx.List().AllLiteral() {
		literal, err := v.visitLiteral(item)
		if err != nil {
			return nil, err
		}
		result = append(result, literal)
	}
	return result, nil
}

//------------------------------------------------------------------------------
// Type
//------------------------------------------------------------------------------

func (v *Visitor) visitType(ctx parser.ITypeContext) (expr.Type, error) {
	var result expr.Type
	err := result.UnmarshalText([]byte(ctx.GetText()))
	return result, err
}

func (v *Visitor) getTreeText(tree antlr.Tree) string {
	return tree.(interface{ GetText() string }).GetText()
}
