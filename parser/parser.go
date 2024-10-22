package parser

import (
	"monkeylang/ast"
	"monkeylang/lexer"
	"monkeylang/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token // useful to peek at the next token to decide what to do next
}

// small helper that advances both curToken and peekToken
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// initialize new instance of a parser - takes in `l` which is a pointer to an instance of a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so that curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // construct the root node of the AST
	program.Statements = []ast.Statement{}

	// iterates over each token in the program until it encounters an EOF, incrementing
	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

// parsers

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// stmt := &ast.LetStatement{Token: p.curToken}
	// todo
	return nil
}
