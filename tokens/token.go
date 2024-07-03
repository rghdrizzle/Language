package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}
var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookUpIndent(ident string) TokenType{
	if tok,ok:= keywords[ident];ok{
		return tok
	}
	return IDENT
}
const (
	ILLEGAL ="ILLEGAL"
	EOF = "EOF"
	IDENT ="IDENT"
	INT ="INT"
	ASSIGN ="="
	PLUS="+"
	COMMA=","
	SEMICOLON=";"
	LPAREN ="{"
	RPAREN="}"
	LBRAC ="("
	RBRAC=")"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	EQ = "=="
	NOT_EQ = "!="
)

