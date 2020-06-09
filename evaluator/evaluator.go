package evaluator

import (
	"github.com/Neeraj-Natu/shifu/ast"
	"github.com/Neeraj-Natu/shifu/object"
)

//Eval is the parent function that calls different evaluators based on what the type of AST node is.
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	//Evaluating Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	//Recursively evaluating each expression
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmnts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmnts {
		result = Eval(statement)
	}
	return result
}
