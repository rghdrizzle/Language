package evaluator

import (
	"rghdrizzle/language/lexer"
	"rghdrizzle/language/objects"
	"rghdrizzle/language/parser"
	"testing"
)

func TestEvalIntegerExpression( t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"5",5},
		{"10",10},
	}
	for _,tt := range tests{
		evaluated:= testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}

}

func testEval(input string) objects.Object{
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}


func testIntegerObject(t *testing.T, obj objects.Object,expected int64)bool{
	result ,ok := obj.(*objects.Integer)
	if !ok{
		t.Errorf("Object is not Integer. got %T{%+v}",obj,obj)
		return false
	}
	if result.Value!=expected{
		t.Errorf("object has wrong value. got=%d, want=%d",result.Value, expected)
		return false
	}
	return true


}
