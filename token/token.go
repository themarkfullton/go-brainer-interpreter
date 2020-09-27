package token

type TokenType string

// Because TokenType is a string, a variety of values can be assigned as TokenTypes
// This allows us to distinguish different types of tokens

// Apparently also easy to debug without a massive amount of boilerplates and helpers
// Ideally, you'd want to use a byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL means a token that is not recognized
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers (variables) and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="

	PLUS = "+"

	MINUS = "-"

	NOT = "!"

	ASTERISK = "*"

	SLASH = "/"

	LT = "<"

	GT = ">"

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

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}


// Checks the keywords map to see if var is keyword
func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}