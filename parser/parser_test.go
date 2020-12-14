package parser

import (
	"testing"
	"github.com/themarkfullton/go-brainer-interpreter/ast"
	"github.com/themarkfullton/go-brainer-interpreter/lexer"
)

func TestLetStatements(t *testing.T){
	input := `let x  5;

let y = 10;
let foobar = 838383;`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t,p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got =%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parsers has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser errpr: %q", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'.got =%q", s.TokenLiteral())

		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {

		t.Errorf("s not *ast.LetStatement. got=%T", s)

		return false

	}

	if letStmt.Name.Value != name {
		t.Errorf("letSmt.Name.Value no '%s'.got %s", name, letStmt.Name.Value)

		return false
	}

	if letStmt.Name.TokenLiteral() != name {

		t.Errorf("letSmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())

		return false
	}

	return true
}