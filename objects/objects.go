package objects

import (
	"fmt"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"

)

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

type Null struct{}

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



