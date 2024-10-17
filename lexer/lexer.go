package lexer

import (
	"monkeylang/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input ^ - points to current char
	readPosition int  // current reading position in input, after current char. We will need this to be able to "peek" further into the input and see what comes next
	ch           byte // current char being examined
}

// init a new lexer
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
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

// helper to init tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
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
	}

	l.readChar()
	return tok
}
