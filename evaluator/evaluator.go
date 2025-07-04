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

func Eval(node ast.Node, env *objects.Environment) objects.Object{
	switch node := node.(type){
	case *ast.Program:
		return evalProgram(node,env)
	case *ast.Identifier:
		return evalIdentifier(node,env)
	case *ast.StringLiteral:
		return &objects.String{Value: node.Value}
	case *ast.IntegerLiteral:
		return &objects.Integer{Value: node.Value}
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &objects.Function{Parameters: params,Body: body, Env: env}
	case *ast.CallExpression:
		function := Eval(node.Function,env)
		if isError(function){
			return function
		}
		args := evalExpressions(node.Arguments,env)
		if len(args)==1 && isError(args[0]){
			return args[0]
		}
		return applyFunction(function, args)
	case *ast.ArrayLiteral:
		elements:= evalExpressions(node.Elements,env)
		if len(elements)==1 &&isError(elements[0]){
			return elements[0]
		}
		return &objects.Array{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(node.Left,env)
		if isError(left){
			return left
		}
		index := Eval(node.Index,env)
		if isError(index){
			return index
		}
		return evalIndexExpression(left,index)

	case *ast.PrefixExpression:
		right := Eval(node.Right,env)
		if isError(right){
			return right
		}
		return evalPrefixExpression(node.Operator,right)
	case *ast.InfixExpression:
		left := Eval(node.Left,env)
		if isError(left){
			return left
		}
		right := Eval(node.Right,env)
		if isError(right){
			return right
		}
		return evalInfixExpression(left,node.Operator,right)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue,env)
		if isError(val){
			return val
		}
		return &objects.RetrunValue{Value: val}
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		
	case *ast.ExpressionStatement:
		return Eval(node.Expression,env)
	case *ast.BlockStatement:
		return evalBlockStatement(node,env )
	case *ast.IfExpression:
		return evalIfExpression(node,env)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.HashLiteral:
		return evalHashLiterals(node, env)
	}
	return nil
}

func evalProgram(program *ast.Program,env *objects.Environment) objects.Object{
	var result objects.Object
	for _,statement := range program.Statements{
		result = Eval(statement,env)
		switch result := result.(type){
		case *objects.RetrunValue:
			return result.Value
		case *objects.Error:
			return result
		}
	}

	return result
}
func evalBlockStatement(block *ast.BlockStatement, env *objects.Environment) objects.Object{
	var result objects.Object
	for _, statements:= range block.Statements{
		result = Eval(statements,env)
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
	case left.Type() == objects.STRING_OBJ && right.Type() == objects.STRING_OBJ:
		return evalStringInfixExpression(op, left, right)
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

func evalIfExpression(ie *ast.IfExpression,env *objects.Environment) objects.Object{
	condition := Eval(ie.Condition,env)
	if isError(condition){
		return condition
	}
	if isTruthy(condition){
		return Eval(ie.Consequence,env)
	}else if ie.Alternative!=nil{
		return Eval(ie.Alternative,env)
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

func evalIdentifier(ident *ast.Identifier,env *objects.Environment) objects.Object{
	obj , ok := env.Get(ident.Value)
	if ok{
		return obj
	}
	if builtin,ok := builtins[ident.Value]; ok{
		return builtin
	}
	return newError("identifier not found: " + ident.Value)
}

func evalExpressions(exp []ast.Expression,env *objects.Environment) []objects.Object{
	var result []objects.Object

	for _, e := range(exp){
		evaluated := Eval(e,env)
		if isError(evaluated){
			return []objects.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(function objects.Object,args []objects.Object) objects.Object{
	switch fn := function.(type){
	case *objects.Function:
		extendedEnv := extendFunctionEnv(fn,args)
		evaluated := Eval(fn.Body,extendedEnv)
		return unwrapReturnValue(evaluated)
	case *objects.BuiltIn:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
	
	
}

func extendFunctionEnv(fn *objects.Function,args []objects.Object) *objects.Environment{
	env := objects.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters{
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj objects.Object) objects.Object{
	if returnValue, ok := obj.(*objects.RetrunValue);ok{
		return returnValue.Value
	}
	return obj
}

func evalStringInfixExpression(op string, left objects.Object, right objects.Object)objects.Object{
	leftValue := left.(*objects.String).Value
	rightValue := right.(*objects.String).Value
	if op=="+"{
		return &objects.String{Value: leftValue+rightValue}
	}else{
		return newError("unknown operator: %s %s %s",left.Type(), op, right.Type())
	}
}

func evalIndexExpression(left objects.Object, index objects.Object) objects.Object{
	switch{
	case left.Type() == objects.ARRAY_OBJ && index.Type() == objects.INTEGER_OBJ:
			return evalArrayIndexExpression(left,index)
	default:
		return newError("index operator not supported: %s",left.Type())
	}
}

func evalArrayIndexExpression(left objects.Object, index objects.Object) objects.Object{
	arrayObj := left.(*objects.Array)
	ind := index.(*objects.Integer).Value
	max := int64(len(arrayObj.Elements)-1)
	if ind<0 || ind > max{
		return NULL
	}
	return arrayObj.Elements[ind]
}
func evalHashLiterals(node *ast.HashLiteral, env *objects.Environment) objects.Object{
	pairs := make(map[objects.HashKey]objects.HashPair)

	for keyNode, valueNode := range node.Pairs{
		key := Eval(keyNode,env)
		if isError(key){
			return key
		}
		hashKey , ok := key.(objects.Hashable)
		if !ok{
			return newError("unusable as hash key: %s", key.Type())
	
		}
		value := Eval(valueNode,env)
		if isError(value){
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = objects.HashPair{Key: key, Value: value}

	}
	return &objects.Hash{Pairs: pairs}
}