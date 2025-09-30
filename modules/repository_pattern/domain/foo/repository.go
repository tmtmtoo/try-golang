package foo

import "repository_pattern/domain/primitives"

type FindFooById interface {
	FindFooById(id primitives.Id) (*Foo, error)
}

type SaveFoo interface {
	SaveFoo(foo *Foo) error
}

type DeleteFoo interface {
	DeleteFoo(id primitives.Id) error
}
