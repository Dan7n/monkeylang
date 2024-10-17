package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input ^ - points to current char
	readPosition int  // current reading position in input, after current char. We will need this to be able to "peek" further into the input and see what comes next
	ch           byte // current char being examined
}

// init a new lexer
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}

func (l *Lexer) readChar() {
	endOfInput := l.readPosition >= len(l.input)
	if endOfInput {
		l.ch = 0 // 0 is ASCI code for NULL
	} else {
		l.ch = l.input[l.readPosition]
	}

	// move to next char
	l.position = l.readPosition
	l.readPosition += 1
}
