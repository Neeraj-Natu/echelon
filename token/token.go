package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers and literals
	VARIABLE = "VAR"     // add, foobar, x, y, ......
	INT      = "INTEGER" // 123424
	STRING   = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	AND      = "&&"
	OR       = "||"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LCBRACE   = "{"
	RCBRACE   = "}"
	LBRACE    = "["
	RBRACE    = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	ELSEIF   = "ELSEIF"
	RETURN   = "RETURN"
	LET      = "LET"
	RANGE    = "RANGE"
	FOR      = "FOR"
	IN       = "IN"
	WHILE    = "WHILE"
	LENGTH   = "LENGTH"
	CONTAINS = "CONTAINS"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSEIF,
	"return": RETURN,
	"for":    FOR,
	"while":  WHILE,
	"range":  RANGE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return VARIABLE
}
