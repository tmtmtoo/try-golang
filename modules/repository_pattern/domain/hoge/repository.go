package hoge

import (
	"repository_pattern/domain/primitives"
)

type LoadHoge interface {
	LoadHoge(id primitives.Id) (Hoge, error)
}

type PersistHoge interface {
	PersistHoge(hoge Hoge) error
}
