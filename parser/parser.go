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

func New(l *lexer.Lexer) *Parser{
  p:= &Parser{l: l}

  p.nextToken()
  p.nextToken()

  return p
}

func (p *Parser) nextToken(){
  p.curToken = p.peekToken
  p.peekToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
  return nil
}
