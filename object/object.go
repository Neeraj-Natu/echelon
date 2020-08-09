package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
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
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
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

//BuiltinFunction accepts zero or more object.Object as arguments and return object.Object
type BuiltinFunction func(args ...Object) Object

//Builtin defines the structure of any builtin function. It's a wrapper over all builtin functions.
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

//Array implements the Object interface. Every ast.ArrayLiteral is converted to this Object.Array
//while evaluating Arrays in the language, reference to this struct is then passed on.
type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

//HashKey is the hashkey that stores the hashed value of the keys for HashLiterals.
//this is used to find if and HashLiteral has a key and returns the value that is stored against that Key in the hashLiteral.
//HashKey can have keys as string, integers or booleans so has differnt HashKey() methods for each of the LiteralType to compare keys.
type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Type of Values in the Hash.Pairs which basically can be anything.
type HashPair struct {
	Key   Object
	Value Object
}

// Hash is the in built datastructure in Shifu language that supports storing the data as Hashes or Maps in other languages.
// It can have Strings, Boolean or Integers as keys which have differnt implementations used to define their own HashKeys
// Those HashKeys are used to ascertain if the keys are similar or different and to compare two hashes or values in hashes.
// Also the HashKeys are used to locate and take out the values from Hashes just like any Map.
// As in any other language Hashes are prone to have hash collision in shifu as well.
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

//This interface is used in evaluator to check if a given object is usable as a hash key
type Hashable interface {
	HashKey() HashKey
}
