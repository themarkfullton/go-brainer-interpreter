package lexer

type Lexer struct {
	input        string
	position     int  //keeps track of current char
	readPosition int  // current position in input
	ch           byte // the char being examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	return l
}

// readChar reads character and advances current position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
