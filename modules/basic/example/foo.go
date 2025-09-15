package example

type Foo struct {
	Name string
}

func NewFoo(name string) *Foo {
	return &Foo{Name: name}
}

func (f *Foo) Greet() string {
	return "Hello, " + f.Name
}
