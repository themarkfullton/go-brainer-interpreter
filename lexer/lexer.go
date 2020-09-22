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
