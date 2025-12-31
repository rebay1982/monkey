package parser

import (
	"testing"

	"monkey/ast"
	"monkey/lexer"
)

func TestBooleanExpression(t *testing.T) {
	input := `
		true;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	testBooleanLiteral(t, stmt.Expression, true)
}
