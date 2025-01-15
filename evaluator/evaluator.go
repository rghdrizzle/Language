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
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(left,node.Operator,right)
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
	case "-":
		return evalMinusPrefixOperatorExpression(right)
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

func evalMinusPrefixOperatorExpression(right objects.Object) objects.Object{
	if right.Type() != objects.INTEGER_OBJ{
		return NULL
	}
	value := right.(*objects.Integer).Value
	return &objects.Integer{Value: -value}
}

func evalInfixExpression(left objects.Object,op string,right objects.Object) objects.Object{
	switch{
	case left.Type()==objects.INTEGER_OBJ && right.Type()==objects.INTEGER_OBJ:
		return evalIntegerInfixExpression(left,op,right)
	case op=="==":
		return nativeBoolToBooleanObject(left==right)
	case op == "!=":
		return nativeBoolToBooleanObject(left!=right)
	default:
		return NULL
	}
	
}

func evalIntegerInfixExpression(left objects.Object,op string,right objects.Object) objects.Object{
	leftValue := left.(*objects.Integer).Value
	rightValue := right.(*objects.Integer).Value
	switch op{
	case "+":
		return &objects.Integer{Value: leftValue+rightValue}
	case "-":
		return &objects.Integer{Value: leftValue-rightValue}
	case "*":
		return &objects.Integer{Value: leftValue*rightValue}
	case "/":
		return &objects.Integer{Value: leftValue/rightValue}
	case ">":
		return nativeBoolToBooleanObject(leftValue>rightValue)
	case "<":
		return nativeBoolToBooleanObject(leftValue<rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue==rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue!=rightValue)
	default:
		return NULL
	
	}
}
