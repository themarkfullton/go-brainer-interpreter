package lexer

import "github.com/themarkfullton/go-brainer-interpreter/token"

type Lexer struct {
	input        string
	position     int  //keeps track of current char
	readPosition int  // current position in input
	ch           byte // the char being examined
}

// This will only support ASCI because ch set to byte
// To use UTF-8, ch would need to be rune and we would need to change the read function

// Go back and figure out how to change to UTF-8 after finishing book

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

// readChar reads character and advances current position
func (l *Lexer) readChar() {
	// Checks if we're at the end of the line, if so, sets to 0/NUL
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// Ensure l.readPosition always points to the next position that will be read and l.position always points to position where we last read.
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch){
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token { Type: tokenType, Literal: string(ch)}
}

// Reads an identifier and advances lexer's position until finds non-letter char
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

// Allows _ to be used in variable names
func isLetter(ch byte) bool{
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}