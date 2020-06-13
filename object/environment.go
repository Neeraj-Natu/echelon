package object

//NewEnclosedEnvironment is used for extending the current environment, it creates a
//new instance of object.Environment with a pointer to the environment it should extend.
//With this we enclose a fresh and empty environment with an existing one.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

//NewEnvironment is just what is used to create a newEnvironment instance.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

//Environment is used to keep track of value by associating them with a name.
//It's just a Hashmap that store the value and name of the variables.
type Environment struct {
	store map[string]Object
	outer *Environment
}

//Get the value of a variable if present in the Environment instance.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

//Set the variable in Environment for later use.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
