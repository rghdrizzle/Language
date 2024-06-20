package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}
var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
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
	SUBT = "-"
	PLUS="+"
	COMMA=","
	SEMICOLON=";"
	LPAREN ="{"
	RPAREN="}"
	LBRAC ="("
	RBRAC=")"

	FUNCTION ="FUNCTION"
	LET ="LET"
)

