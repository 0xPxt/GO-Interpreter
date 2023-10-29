package lexer

import (
	"source/token"
)

type Lexer struct {
	input 			string
	position 		int 	// Current position in input
	readPosition 	int 	// Next position in input
	ch 				byte
}

var keywords = map[string]token.TokenType { 
	"ILLEGAL" 	: token.ILLEGAL,
	"EOF" 		: token.EOF,
	"IDENT" 	: token.IDENT,
	"INT" 		: token.INT,
	"," 		: token.COMMA,
	";" 		: token.SEMICOLON,
	"(" 		: token.LPAREN,
	")" 		: token.RPAREN,
	"{" 		: token.LBRACE,
	"}" 		: token.RBRACE,
	"FUNCTION" 	: token.FUNCTION,
	"LET" 		: token.LET,
	"=" 		: token.ASSIGN,
	"+"  		: token.PLUS,
	"-"  		: token.MINUS,
	"!"  		: token.BANG,
	"=="		: token.EQUALS,
	"!="		: token.NOT_EQUALS,
	"*"  		: token.ASTERISK,
	"/"  		: token.SLASH,
	"<"  		: token.LT,
	">" 		: token.GT,
} 


// Create instance of Lexer structure and return a pointer to it
func New(input string) *Lexer {
	lexer := &Lexer{input: input}

	lexer.readChar();

	// Returns pointer to Lexer struct initialized with input field to @param input
	return lexer
}

func (lexer *Lexer)NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	if lexer.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
	}
	if isLetter(lexer.ch) {
		literal := lexer.readIdentifier()
		tok = newToken(token.LookupIdent(literal), literal);
		return tok
	} else if isDigit(lexer.ch) {
		tok = newToken(token.INT, lexer.readNumber());
		return tok
	} else if(lexer.ch == '=') {
		// Special case for "=="
		if lexer.lookAhead() == '=' {
			// Confirm read
			lexer.readChar();
			tok = newToken(token.EQUALS, "==")
		}
	} else if (lexer.ch == '!') {
		// Special case for "!="
		if lexer.lookAhead() == '=' {
			// Confirm read
			lexer.readChar();
			tok = newToken(token.NOT_EQUALS, "!=")
		}
	}

	if tok.Type == "" && tok.Literal == "" {
		tok = newToken(keywords[string(lexer.ch)], string(lexer.ch))
	}


	if tok.Type == "" && tok.Literal == "" {
		tok = newToken(token.ILLEGAL, string(lexer.ch))
	}

	lexer.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

// Assign next character and advance position in input
func (lexer *Lexer)readChar() {
	if (lexer.readPosition >= len(lexer.input)) {
		lexer.ch = 0 // This is a character => NULL (ASCII 0)
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool { 
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isDigit(ch byte) bool { 
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar() 
	}
}

func (lexer *Lexer) lookAhead() byte {
	if lexer.readPosition > len(lexer.input) { 
		return 0
	}

	return lexer.input[lexer.readPosition]
}
