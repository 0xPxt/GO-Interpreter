package ast

import (
	"source/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
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

type IntegerLiteral struct {
	Token token.Token
	Value int64
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

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	}

	return ""
}

func (prog *Program) String() string {
	var buff bytes.Buffer

	for _, statement := range prog.Statements {
		buff.WriteString(statement.String())
	}

	return buff.String()
}

func (letSt *LetStatement)StatementNode() {

}

func (letSt *LetStatement)TokenLiteral() string {
	return letSt.Token.Literal
}

func (letSt *LetStatement)String() string {
	var buff bytes.Buffer

	buff.WriteString(letSt.TokenLiteral() + " ")
	buff.WriteString(letSt.Name.String())

	if letSt.Value != nil {
		buff.WriteString(" = ")
		buff.WriteString(letSt.Value.String())
	}

	buff.WriteString(";")

	return buff.String()
}

func (id *Identifier)ExpressionNode() {

}

func (id *Identifier)TokenLiteral() string {
	return id.Token.Literal
}

func (id *Identifier)String() string {
	return id.Value
}

func (intL *IntegerLiteral)ExpressionNode() {

}

func (intL *IntegerLiteral)TokenLiteral() string {
	return intL.Token.Literal
}

func (intL *IntegerLiteral)String() string {
	return intL.Token.Literal
}

func (retSt *ReturnStatement)StatementNode() {

}

func (retSt *ReturnStatement)ExpressionNode() {

}

func (retSt *ReturnStatement)TokenLiteral() string {
	return retSt.Token.Literal
}

func (retSt *ReturnStatement)String() string {
	var buff bytes.Buffer

	buff.WriteString(retSt.TokenLiteral() + " ")

	if retSt.ReturnValue != nil {
		buff.WriteString(retSt.ReturnValue.String())
	}

	return buff.String()
}

func (expSt *ExpressionStatement)StatementNode() {

}

func (expSt *ExpressionStatement)ExpressionNode() {

}

func (expSt *ExpressionStatement)TokenLiteral() string {
	return expSt.Token.Literal
}

func (expSt *ExpressionStatement)String() string {
	if expSt.Expression != nil {
		return expSt.Expression.String()
	}

	return ""
}
