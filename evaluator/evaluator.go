package evaluator


import (
	"rghdrizzle/language/ast"
	"rghdrizzle/language/objects"

)
var (
	TRUE = &objects.Boolean{Value: true}
	FALSE = &objects.Boolean{Value: false}
	NULL = &objects.Null{}
)

func Eval(node ast.Node) objects.Object{
	switch node := node.(type){
	case *ast.Program:
		return evalStatement(node.Statements)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator,right)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.IntegerLiteral:
		return &objects.Integer{Value: node.Value}
	}
	return nil
}

func evalStatement(statements []ast.Statement) objects.Object{
	var result objects.Object
	for _,statement := range statements{
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBooleanObject(value bool) objects.Object{
	if value==true{
		return TRUE
	}else{
		return FALSE
	}
}

func evalPrefixExpression(operator string,right objects.Object) objects.Object{
	switch operator{
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
	
}

func evalBangOperatorExpression(right objects.Object)objects.Object{
	switch right{
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}

}