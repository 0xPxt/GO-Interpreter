package token

type TokenType string

type Token struct {
	Type TokenType

	Literal string
}


// Define all possible TokenTypes as constants

const (
	ILLEGAL 	= "ILLEGAL"	
	EOF 		= "EOF"

	IDENT 		= "IDENT"
	INT 		= "INT"

	COMMA 		= ","
	SEMICOLON 	= ";"

	LPAREN 		= "("
	RPAREN 		= ")"

	LBRACE 		= "{"
	RBRACE 		= "}"

	ASSIGN   	= "="
	PLUS     	= "+"
	MINUS    	= "-"
	BANG     	= "!"

	EQUALS 		= "=="
	NOT_EQUALS 	= "!="

	ASTERISK 	= "*"
	SLASH    	= "/"
	LT 			= "<"
	GT 			= ">"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	TRUE     = "TRUE"
	FALSE    = "FALSE"

	IF       = "IF"
	ELSE     = "ELSE"

	RETURN   = "RETURN"
)

var keywords = map[string]TokenType {
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
