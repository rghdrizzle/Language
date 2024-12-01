package ast

import (
	"bytes"
	"rghdrizzle/language/tokens"
	"strings"
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

type Boolean struct {
	Token token.Token
	Value bool
}
func (b *Boolean) ExpressionNode() {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string { return b.Token.Literal }


type IfExpression struct{
	Token token.Token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}
func (Fe *IfExpression) ExpressionNode(){}
func (Fe *IfExpression) TokenLiteral()string{ return Fe.Token.Literal}
func(Fe *IfExpression) String() string{ 
	var buf bytes.Buffer

	buf.WriteString("if")
	buf.WriteString(Fe.Condition.String())
	buf.WriteString("")
	buf.WriteString(Fe.Consequence.String())
	if(Fe.Alternative!=nil){
		buf.WriteString("else")
		buf.WriteString(Fe.Alternative.String())
	}
	return buf.String()
}


type BlockStatement struct{
	Token token.Token
	Statements []Statement
}
func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}


type FunctionLiteral struct{
	Token token.Token
	Parameters []*Identifier
	Body *BlockStatement
}
func (fl *FunctionLiteral) ExpressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string{
	var out  bytes.Buffer

	params := []string{}
	for _ , p := range fl.Parameters{
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()

}

type CallExpression struct{
	Token token.Token
	Function Expression
	Arguments []Expression
}
func (ce *CallExpression) ExpressionNode(){}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string{
	var out bytes.Buffer

	args := []string{}
	for _,a := range ce.Arguments{
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
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