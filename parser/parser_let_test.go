package parser

import (
	"testing"

	"monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// TODO: Need to implement the identifier value parsing
//func TestLetStatements_EXT(t *testing.T) {
//	tests := []struct {
//		input              string
//		expectedIdentifier string
//		expectedValue      any
//	}{
//		{"let x = 5;", "x", 5},
//		{"let y = true;", "y", true},
//		{"let foobar = y;", "foobar", "y"},
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
//				t.Fatalf("program.Statements does not contain 1 statements. got=%d",
//					len(program.Statements))
//			}
//
//			stmt := program.Statements[0]
//			if !testLetStatement(t, stmt, tt.expectedIdentifier) {
//				return
//			}
//
//			val := stmt.(*ast.LetStatement).Value
//			fmt.Printf("val %v", val)
//			if !testLiteralExpression(t, val, tt.expectedValue) {
//				return
//			}
//		})
//	}
//}
