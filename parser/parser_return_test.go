package parser

import (
	"testing"

	"monkey/ast"
	"monkey/lexer"
)

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

// TODO: Need to implement the return value parsing
//func TestReturnStatements_EXT(t *testing.T) {
//	tests := []struct {
//		input         string
//		expectedValue any
//	}{
//		{"return 5;", 5},
//		{"return true;", true},
//		{"return foobar;", "foobar"},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.input, func(t *testing.T) {
//			l := lexer.New(tt.input)
//			p := New(l)
//			program := p.ParseProgram()
//			checkParserErrors(t, p)
//
//			if len(program.Statements) != 1 {
//				t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
//			}
//
//			stmt := program.Statements[0]
//			returnStmt, ok := stmt.(*ast.ReturnStatement)
//			if !ok {
//				t.Fatalf("stmt not *ast.ReturnStatement. got=%T", stmt)
//			}
//			if returnStmt.TokenLiteral() != "return" {
//				t.Fatalf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
//			}
//			if testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue) {
//				return
//			}
//		})
//	}
//}
