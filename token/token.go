package token

type TokenType string

// Because TokenType is a string, a variety of values can be assigned as TokenTypes
// This allows us to distringuish different types of tokens

// Apparently also easy to debug without a massive amount of boilerplates and helpers
// Ideally, you'd want to use a byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers (variables) and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="

	PLUS = "+"

	// Delimiters
	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"

	LET = "LET"
)
