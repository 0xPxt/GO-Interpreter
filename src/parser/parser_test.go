package parser

import (
	"testing"
	"source/ast"
	"source/lexer"
)

func TestLetStatement(t *testing.T) {
	input :=
	`
	let x 5;
	let = 10;
	let 838383;
	`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()

	checkParseErrors(t, parser.Errors());

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if (len(program.Statements) > 3) {
		t.Fatalf("got : %d statements in program, it should have been 3", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, errors []string) {
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, message:= range errors {
		t.Errorf("Parser error : %q", message)
	}

	t.FailNow()
}

func testLetStatement( t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral of statement is not 'let', got : %q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)

	if !ok {
		t.Errorf("statement not *ast.LetStatement, got : %T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got : %s", name, letStatement.Name.Value)
		return false
	}

	return true
}