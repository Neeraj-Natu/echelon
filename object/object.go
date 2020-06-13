package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Neeraj-Natu/shifu/ast"
)

type ObjectType string

//Object is the high level interface that represents every value in the SHIFU language,
//this is used in the evaluation step where AST node is converted to this Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

//Integer implements Object interface. Every ast.IntegerLiteral is converted to this Object.Integer
//When evaluating the language, the reference to this struct is then passed around.
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

//Boolean implements the Object interface. Every ast.BooleanLiteral is converted to this Object.Boolean
// When evaluating the language, the reference to this struct is then passed around.
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

//String implements the Object interface. Every ast.StringLiteral is converted to this Object.String
//while evaluating the language, the reference to this struct si then passed around.
type String struct {
	Value string
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

//Null implements the Object interface. This represents the abscene of value.
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

//ReturnValue is a wrapper around the object that is returned with 'return' call.
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

//Error is the object that is returned when an invalid syntax is used while writing programs in language.
//this is different from exception handling that is done within a program written in the language.
type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

//Function implements the Object interface. Every ast.FunctionLiteral is converted to this Object.Function
//while evaluating functions in the language, reference to this struct is then passed on.
//Also any variables are all stored in the environment
type Function struct {
	Parameters []*ast.Variable
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
