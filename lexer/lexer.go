package lexer

import (
	"monkeylang/token"
)

const (
	TAB_CHAR      = '\t'
	NEW_LINE_CHAR = '\n'
	CR_CHAR       = '\r'
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhitespace()

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
		// identifiers and keywords
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isInt(l.ch) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// helper to init a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// useful util to skip over every whitespace char since it doesn't have any meaning in our lang
func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == TAB_CHAR || l.ch == NEW_LINE_CHAR || l.ch == CR_CHAR {
		l.readChar()
	}
}

// reads an identifier and advances the lexer's position until it encounters a non-letter character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readInt() string {
	position := l.position
	for isInt(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isInt(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
