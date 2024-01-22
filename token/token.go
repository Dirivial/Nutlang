package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT   = "IDENT" // add, foobar, x, y, ...
	INT     = "INT"
	FLOAT   = "FLOAT"
	COMMENT = "COMMENT"

	// Operators
	BIND     = ":="
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"

	LT  = "<"
	LTE = "<="
	GT  = ">"
	GTE = ">="

	EQ     = "=="
	NOT_EQ = "!="
	AND    = "&&"
	OR     = "||"

	BITWISEAND = "&"
	BITWISEOR  = "|"

	// Delimiters
	COMMA     = ","
	DOT       = "."
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	ELSE     = "ELSE"
	IF       = "IF"
	STRING   = "STRING"
	FOR      = "FOR"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
