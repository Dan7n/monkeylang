package ast

import "monkeylang/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// LetStatement satisfies the Node interface by returning a statementNode() method
func (ls *LetStatement) statementNode() {}

// LetStatement satisfies the Node interface by returning a TokenLiteral()
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
