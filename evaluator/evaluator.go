package evaluator


import (
	"rghdrizzle/language/ast"
	"rghdrizzle/language/objects"

)

func Eval(node ast.Node) objects.Object{
	switch node := node.(type){
	case *ast.Program:
		return evalStatement(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
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