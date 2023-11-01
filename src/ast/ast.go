package ast

import (
	"source/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

// Root node of the AST
type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	}

	return ""
}

func (letSt *LetStatement)StatementNode() {

}

func (letSt *LetStatement)TokenLiteral() string {
	return letSt.Token.Literal
}

func (id *Identifier)ExpressionNode() {

}

func (id *Identifier)TokenLiteral() string {
	return id.Token.Literal
}

func (retSt *ReturnStatement)StatementNode() {

}

func (retSt *ReturnStatement)ExpressionNode() {

}

func (retSt *ReturnStatement)TokenLiteral() string {
	return retSt.Token.Literal
}
