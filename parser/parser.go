package parser


import(
	"rghdrizzle/language/ast"
	"rghdrizzle/language/tokens"
	"rghdrizzle/language/lexer"
)


type Parser struct{
	l *lexer.Lexer
	

	curToken token.Token
	peekToken token.Token
}