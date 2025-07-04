package evaluator

import (
	"rghdrizzle/language/objects"
)

var builtins = map[string]*objects.BuiltIn{
	"len": &objects.BuiltIn{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *objects.Array:
				return &objects.Integer{Value: int64(len(arg.Elements))}
			case *objects.String:
				return &objects.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &objects.BuiltIn{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},
	"last": &objects.BuiltIn{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"rest": &objects.BuiltIn{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]objects.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])

				return &objects.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": &objects.BuiltIn{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)
			newElements := make([]objects.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &objects.Array{Elements: newElements}
		},
	},
	// "punch": &objects.BuiltIn{
	// 	Fn: func(args ...objects.Object)objects.Object{
			
	// 	},
	// },
}
