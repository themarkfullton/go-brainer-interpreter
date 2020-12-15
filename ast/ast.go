package ast

import (
	"github.com/themarkfullton/go-brainer-interpreter/token"
	"bytes"
	)

type Node interface {
	TokenLiteral() string
	String() string
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Statement interface {
	Node

	statementNode()
}

type Expression interface {
	Node

	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string{
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}