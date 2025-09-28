package foo

type FindFooById interface {
	FindFooById(id Id) (*Foo, error)
}

type SaveFoo interface {
	SaveFoo(foo *Foo) error
}

type DeleteFoo interface {
	DeleteFoo(id Id) error
}
