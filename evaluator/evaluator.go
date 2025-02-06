package evaluator

import (
	"fmt"
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
		return evalProgram(node)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		if isError(right){
			return right
		}
		return evalPrefixExpression(node.Operator,right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left){
			return left
		}
		right := Eval(node.Right)
		if isError(right){
			return right
		}
		return evalInfixExpression(left,node.Operator,right)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		if isError(val){
			return val
		}
		return &objects.RetrunValue{Value: val}
	case *ast.LetStatement:
		idt := node.Name
		val := Eval(node.Value)
		if isError(val) {
			return val
		}			
		return evalLetStatement(idt,val) 
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.IntegerLiteral:
		return &objects.Integer{Value: node.Value}
	}
	return nil
}

func evalProgram(program *ast.Program) objects.Object{
	var result objects.Object
	for _,statement := range program.Statements{
		result = Eval(statement)
		switch result := result.(type){
		case *objects.RetrunValue:
			return result.Value
		case *objects.Error:
			return result
		}
	}

	return result
}
func evalBlockStatement(block *ast.BlockStatement) objects.Object{
	var result objects.Object
	for _, statements:= range block.Statements{
		result = Eval(statements)
		if result!=nil {
			rt := result.Type() 
			if rt == objects.RETURN_VALUE_OBJ || rt == objects.ERROR_OBJ{
				return result
			}
		}
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
		return newError("unknown operator: %s%s", operator, right.Type())
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
		return newError("unknown operator: -%s", right.Type())
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",left.Type(), op, right.Type())
	default:
		return newError("unknown operator: %s %s %s",left.Type(), op, right.Type())
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
		return newError("unknown operator: %s %s %s",left.Type(), op, right.Type())
	
	}
}

func evalIfExpression(ie *ast.IfExpression) objects.Object{
	condition := Eval(ie.Condition)
	if isError(condition){
		return condition
	}
	if isTruthy(condition){
		return Eval(ie.Consequence)
	}else if ie.Alternative!=nil{
		return Eval(ie.Alternative)
	}else{
		return NULL
	}

}

func isTruthy(obj objects.Object)bool{
	switch obj{
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string ,a ...interface{}) *objects.Error{
	return &objects.Error{Message: fmt.Sprintf(format,a...)}
}
func isError(obj objects.Object)bool{
	if obj!=nil{
		return obj.Type() == objects.ERROR_OBJ
	}
	return false
}
