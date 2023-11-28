package parser

import (
	"testing"
	"source/ast"
	"source/lexer"
)

func TestLetStatement(t *testing.T) {
	input :=
	`
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()

	checkParseErrors(t, parser.Errors());

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if (len(program.Statements) != 3) {
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

func TestReturnStatement(t *testing.T) {
	input :=
	`
	return 5;
	return 10;
	return 993322;
	`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()

	checkParseErrors(t, parser.Errors());

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if (len(program.Statements) != 3) {
		t.Fatalf("got : %d statements in program, it should have been 3", len(program.Statements))
	}

	for _, statement := range program.Statements {
		if !testReturnStatement(t, statement) {
			return
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(t, parser.Errors())

	if len(program.Statements) != 1 {
		t.Fatalf("program has statements != 1, got '%d'", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got '%T'", program.Statements[0])
	}

	ident, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("program.Statements[0].Expression is not ast.Identifier, got '%T'", statement.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not 'foobar', got '%s'", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not 'foobar', got '%s'", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t* testing.T) {
	input := "5;"

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(t, parser.Errors())

	if len(program.Statements) != 1 {
		t.Fatalf("program has statements != 1, got '%d'", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got '%T'", program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral, got '%T'", statement.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("literal.Value not 5, got '%d'", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Fatalf("literal.TokenLiteral() not 5, got '%s'", literal.TokenLiteral())
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

func testReturnStatement( t *testing.T, statement ast.Statement) bool {
	if statement.TokenLiteral() != "return" {
		t.Errorf("TokenLiteral of statement is not 'return', got : %q", statement.TokenLiteral())
		return false
	}

	_, ok := statement.(*ast.ReturnStatement)

	if !ok {
		t.Errorf("statement not *ast.LetStatement, got : %T", statement)
		return false
	}

	return true
}
