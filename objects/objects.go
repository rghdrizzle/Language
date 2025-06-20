package objects

import (
	"fmt"
	"rghdrizzle/language/ast"
	"strings"
	"bytes"
	"hash/fnv"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ = "STRING"
	BUILTIN_OBJ = "BUILTIN"
	ARRAY_OBJ = "ARRAY"

)
type BuiltInFunction func(args ...Object) Object

type ObjectType string

type Object interface{
	Type() ObjectType
	Inspect() string
}

type Integer struct{
	Value int64
}

type Boolean struct{
	Value bool
}

type RetrunValue struct{
	Value Object
}

type Error struct{
	Message string
}
type Function struct{
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}
type String struct{
	Value string
}
type Array struct{
	Elements []Object
}

type Null struct{}

type BuiltIn struct{
	Fn BuiltInFunction
}
type HashKey struct{
	Type ObjectType
	Value uint64
}

func (i *Integer) Inspect() string{
	return fmt.Sprintf("%d",i.Value)
}

func (i *Integer) Type() ObjectType{
	return INTEGER_OBJ
}

func (b *Boolean) Inspect() string{
	return fmt.Sprintf("%t",b.Value)
}

func (b *Boolean) Type() ObjectType{
	return BOOLEAN_OBJ
}
func (n *Null) Inspect() string{
	return fmt.Sprintf("null")
}

func (n *Null) Type() ObjectType{
	return NULL_OBJ
}

func (rv *RetrunValue) Type() ObjectType{
	return RETURN_VALUE_OBJ
}

func (rv *RetrunValue) Inspect() string{
	return rv.Value.Inspect()
}

func (e *Error) Type() ObjectType{
	return ERROR_OBJ
}
func (e *Error) Inspect() string{
	return "ERROR:"+e.Message
}

func (f *Function) Type() ObjectType{ 
	return FUNCTION_OBJ 
}
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (s *String) Type() ObjectType{
	return STRING_OBJ
}

func (s *String) Inspect() string{
	return s.Value
}

func (b *BuiltIn) Type() ObjectType{
	return BUILTIN_OBJ
}

func (b *BuiltIn) Inspect() string{
	return "builtin function"
}
func (ao *Array) Type() ObjectType{
	return ARRAY_OBJ
}

func (ao *Array) Inspect() string{
	var out bytes.Buffer

	elements :=[]string{}
	for _,e := range ao.Elements{
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements,", "))
	out.WriteString("]")

	return out.String()
}

func (b *Boolean) HashKey() HashKey{
	var value uint64
	 if b.Value {
		value =1
	 }else{
		value =0
	 }
	 return HashKey{Type: b.Type(),Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
func (s *String) HashKey() HashKey{
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(),Value: h.Sum64()}
}