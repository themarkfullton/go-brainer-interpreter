package parser

import (
	"github.com/themarkfullton/go-brainer-interpreter/ast"
	"github.com/themarkfullton/go-brainer-interpreter/lexer"
	"github.com/themarkfullton/go-brainer-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token

	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{1:1}

	p.nextToken()

	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken

	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}