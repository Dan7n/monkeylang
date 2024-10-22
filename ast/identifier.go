package ast

import "monkeylang/token"

// implements the `Expression` interface - holds the IDENT token and its value
type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) expressionNode() {}
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}
