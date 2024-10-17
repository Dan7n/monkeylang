package token

type TokenType string // using string for easy debugging without a bunch of boilerplate and helpers fns: just print out the string to stdout

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // chars/tokens we don't support or know about
	EOF     = "EOF"     // end of file - tells our parser that it can stop

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)