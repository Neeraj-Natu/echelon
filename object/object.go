package object

import "fmt"

type ObjectType string

//Object is the high level interface that represents every value in the SHIFU language,
//this is used in the evaluation step where AST node is converted to this Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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

//Null implements the Object interface. This represents the abscene of value.
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }
