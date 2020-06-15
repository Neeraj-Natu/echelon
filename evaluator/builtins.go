package evaluator

import (
	"github.com/Neeraj-Natu/shifu/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'first' must be an ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return newError("an array has no elements!")
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'last' must be an ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return newError("an array has no elements!")
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, expected=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'push' must be an ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"pop": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'pop' must be an ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to 'pop' must be an INTEGER, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			idx := args[1].(*object.Integer)
			index := idx.Value
			if index < 0 || index >= int64(length) {
				return newError("Index to pop from is out of bounds!")
			}
			arr.Elements = append(arr.Elements[:index], arr.Elements[index:]...)
			return arr
		},
	},
}
