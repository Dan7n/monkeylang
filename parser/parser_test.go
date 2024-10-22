package parser

import (
	"monkeylang/ast"
	"monkeylang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
  let y = 10;
  let foobar = 838383;
  `

	lxr := lexer.New(input)
	p := New(lxr)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 elements - received '%d'", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for idx, tst := range tests {
		statement := program.Statements[idx]
		if !testStatement(t, statement, tst.expectedIdentifier) {
			return
		}
	}
}

func testStatement(t *testing.T, statement ast.Statement, expectedName string) bool {
	const SUPPORTED_TOKEN_LITERAL = "let"
	if statement.TokenLiteral() != SUPPORTED_TOKEN_LITERAL {
		t.Errorf("s.TokenLiteral is not '%q', received '%q'", SUPPORTED_TOKEN_LITERAL, statement.TokenLiteral())
		return false
	}

	letStmnt, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not *ast.LetStatement - received '%T'", statement)
		return false
	}

	if letStmnt.Name.Value != expectedName {
		t.Errorf("statement.Name.Value not '%s' - received '%s'", expectedName, letStmnt.Name.Value)
		return false
	}

	if letStmnt.Name.TokenLiteral() != expectedName {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s' - received '%s'", expectedName, letStmnt.Name.TokenLiteral())
		return false
	}

	return true
}
