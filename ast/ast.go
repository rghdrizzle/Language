package ast
import "rghdrizzle/language/tokens"

type Node interface{
	TokenLiteral() string
}

type Statement interface{
	Node // embedding the Node interface , so the statement interface must also implement the methods of Node interface
	StatementNode()
}

type Expression interface{
	Node // embedding the Node interface , so the statement interface must also implement the methods of Node interface
	ExpressionNode()
}

type Program struct{
	Statements []Statement // A slice of statement objects
}

type Identifier struct{
	Token token.Token
	Value string
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct{
	Token token.Token // token.Let 
	Name *Identifier
	Value Expression
}
func (l *LetStatement) StatementNode(){}
func (l *LetStatement) TokenLiteral() string{ return l.Token.Literal}

func (p *Program) TokenLiteral() string{
	if len(p.Statements)>0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}
