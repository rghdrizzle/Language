package ast

import (
	"bytes"
	"rghdrizzle/language/tokens"
)

type Node interface{
	TokenLiteral() string
	String() string
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

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct{
	Token token.Token // token.Let 
	Name *Identifier
	Value Expression
}
func (l *LetStatement) StatementNode(){}
func (l *LetStatement) TokenLiteral() string{ return l.Token.Literal}

type ReturnStatement struct {
	Token token.Token //token.Return
	ReturnValue Expression
	}
func (rs *ReturnStatement) StatementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct{
	Token token.Token 
	Expression Expression
}
func (es *ExpressionStatement) StatementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type IntegerLiteral struct{
	Token token.Token
	Value int64
}
func (il *IntegerLiteral) ExpressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

type PrefixExpression struct{
	Token token.Token
	Operator string
	Right Expression
}
func (pe *PrefixExpression) ExpressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}


type InfixExpression struct{
	Token token.Token
	Left Expression
	Operator string
	Right Expression
}
func (oe *InfixExpression) ExpressionNode(){}
func (oe *InfixExpression) TokenLiteral() string{ return oe.Token.Literal}
func (oe *InfixExpression) String() string{
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements)>0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}
func (p *Program) String() string{
	var out bytes.Buffer
	for _,s := range p.Statements{
		out.WriteString(s.String())
	}
	return out.String()
}
func (ls *LetStatement) String() string {
    var out bytes.Buffer
    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")
    if ls.Value != nil {
    out.WriteString(ls.Value.String())
    }
    out.WriteString(";")
    return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
	out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
	return es.Expression.String()
	}
	return ""
}
func (i *Identifier) String() string { 
	return i.Value 
}