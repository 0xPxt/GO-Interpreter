package parser

import(
	"fmt"
	"source/ast"
	"source/lexer"
	"source/token"
)

type Parser struct {
	lexer *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
	errors []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}

	// Read two tokens to set currentToken and peekToken
	parser.nextToken()
	parser.nextToken()

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
	default:
		return nil
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
	error := fmt.Sprintf("[Peek Error] Expected next token to be : %s, got %s instead.",
		tt, parser.peekToken.Type);

	parser.errors = append(parser.errors, error)
}
