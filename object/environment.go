package object

//NewEnvironment is just what is used to create a newEnvironment instance.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

//Environment is used to keep track of value by associating them with a name.
//It's just a Hashmap that store the value and name of the variables.
type Environment struct {
	store map[string]Object
}

//Get the value of a variable if present in the Environment instance.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

//Set the variable in Environment for later use.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
