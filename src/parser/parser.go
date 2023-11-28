package parser

import(
	"fmt"
	"source/ast"
	"source/lexer"
	"source/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	lexer *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken token.Token

	prefixParseMap map[token.TokenType]prefixParseFunc
	infixParseMap map[token.TokenType]infixParseFunc
}

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc func(ast.Expression) ast.Expression
)

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	parser.prefixParseMap[tokenType] = fn
}

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	parser.infixParseMap[tokenType] = fn
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}

	// Read two tokens to set currentToken and peekToken
	parser.nextToken()
	parser.nextToken()

	parser.prefixParseMap = make(map[token.TokenType]prefixParseFunc)
	parser.registerPrefix(token.IDENT, parser.parseIdentifier)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)

	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser)ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parser.isCurrentToken(token.EOF) {
		statement := parser.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENT) {
		// Next token should be the variable name
		return nil
	}

	letStatement.Name = &ast.Identifier{Token: parser.currentToken, Value:parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		// Next token should be '='
		return nil
	}

	for !parser.isCurrentToken(token.SEMICOLON) {
		// TODO: Parse expression instead of skipping it
		parser.nextToken()
	}

	return letStatement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	for !parser.isCurrentToken(token.SEMICOLON) {
		// TODO: Parse expression instead of skipping it
		parser.nextToken()
	}

	return returnStatement
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: parser.currentToken}

	statement.Expression = parser.parseExpression(LOWEST)

	if parser.isPeekToken(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseMap[parser.currentToken.Type]

	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	intL := &ast.IntegerLiteral{Token: parser.currentToken}

	value, error := strconv.ParseInt(parser.currentToken.Literal, 0 , 64)

	if error != nil {
		message := fmt.Sprintf("Could not parse %q as integer", parser.currentToken.Literal)
		parser.errors = append(parser.errors, message)
		return nil
	}
	intL.Value = value
	return intL
}

func (parser *Parser) isCurrentToken(tt token.TokenType) bool {
	return parser.currentToken.Type == tt
}

func (parser *Parser) isPeekToken(tt token.TokenType) bool {
	return parser.peekToken.Type == tt	
}

// Assert peeked token
func (parser *Parser) expectPeek(tt token.TokenType) bool {
	if parser.isPeekToken(tt) {
		parser.nextToken()
		return true
	} else {
		parser.addPeekError(tt)
		return false
	}
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) addPeekError(tt token.TokenType) {
	error := fmt.Sprintf("[Peek Error] Expected next token to be '%s', got '%s' instead.",
		tt, parser.peekToken.Type);

	parser.errors = append(parser.errors, error)
}
