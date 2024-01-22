package evaluator

import (
	"Nutlang/object"
	"fmt"
	mrand "math/rand"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},

	"min": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				switch arg2 := args[1].(type) {
				case *object.Integer:
					return &object.Integer{Value: min(arg.Value, arg2.Value)}
				default:
					return newError("argument 2 to `max` must be INTEGER, got %s",
						args[0].Type())
				}
			default:
				return newError("argument 1 to `max` must be INTEGER, got %s",
					args[0].Type())
			}
		},
	},
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				switch arg2 := args[1].(type) {
				case *object.Integer:
					return &object.Integer{Value: max(arg.Value, arg2.Value)}
				default:
					return newError("argument 2 to `max` must be INTEGER, got %s",
						args[0].Type())
				}
			default:
				return newError("argument 1 to `max` must be INTEGER, got %s",
					args[0].Type())
			}
		},
	},
	"rand": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return &object.Integer{Value: int64(mrand.Int())}
			} else if len(args) == 1 {
				if args[0].Type() == object.INTEGER_OBJ {
					arg := args[0].(*object.Integer).Value
					return &object.Integer{Value: int64(mrand.Int63n(arg))}
				}
				return newError("argument to `first` must be INTEGER, got %s",
					args[0].Type())
			} else {
				return newError("wrong number of arguments. got=%d, want=0 or 1",
					len(args))
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"pop": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length == 0 {
				return &object.Array{Elements: []object.Object{}}
			}
			if length > 0 {

				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[0:length-1])

				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},

	"unshift": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument 1 to `unshift` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length == 0 {
				return &object.Array{Elements: []object.Object{}}
			}
			if length > 0 {
				switch arg2 := args[1].(type) {
				case *object.Integer:

					newElements := make([]object.Object, length+1)
					copy(newElements[1:], arr.Elements)
					newElements[0] = arg2

					return &object.Array{Elements: newElements}
				case *object.Float:

					newElements := make([]object.Object, length+1)
					copy(newElements[1:], arr.Elements)
					newElements[0] = arg2

					return &object.Array{Elements: newElements}
				case *object.Boolean:
					newElements := make([]object.Object, length+1)
					copy(newElements[1:], arr.Elements)
					newElements[0] = arg2

					return &object.Array{Elements: newElements}
				case *object.String:

					newElements := make([]object.Object, length+1)
					copy(newElements[1:], arr.Elements)
					newElements[0] = arg2

					return &object.Array{Elements: newElements}
				case *object.Array:

					length2 := len(arg2.Elements)
					if length2 == 0 {
						return arr
					}
					newElements := make([]object.Object, length+length2)
					copy(newElements[length2:], arr.Elements)
					copy(newElements[:length2], arg2.Elements)

					return &object.Array{Elements: newElements}
				default:
					return newError("argument 2 to `unshift` not supported, got %s",
						args[1].Type())
				}
			}
			return NULL
		},
	},
	"shift": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `shift` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length == 0 {
				return &object.Array{Elements: []object.Object{}}
			}
			if length > 0 {

				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])

				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
