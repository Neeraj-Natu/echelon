package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Indentifiers and literals
	VARIABLE = "VAR" // add, foobar, x, y, ......
	INT = "INT" // 123424
	FLOAT = "FLOAT" // 1231234.31251512

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	AND = "&"
	OR = "||"
	
	EQ = "=="
	NOT_EQ = "!="
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	COLON = ":"
	LPAREN = "("
	RPAREN = ")"
	LCBRACE = "{"
	RCBRACE = "}"
	LBRACE = "["
	RBRACE = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return VARIABLE
}