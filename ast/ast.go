package ast

// every node in our AST has to provide a TokenLiteral() method that returns the literal value of the token it's associated with
type Node interface {
	TokenLiteral() string // used for debugging and testing
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// `Program` is the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
