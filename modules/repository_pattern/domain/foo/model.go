package foo

import (
	"fmt"
	"repository_pattern/domain/primitives"
)

type Bar struct {
	Id    primitives.Id
	Value primitives.Text
}

func NewBar(id []byte, value string) (*Bar, error) {
	parsedId, err := primitives.ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	bar := &Bar{
		Id:    parsedId,
		Value: primitives.NewText(value),
	}

	return bar, nil
}

type Baz struct {
	Id    primitives.Id
	Value primitives.Text
}

func NewBaz(id []byte, value string) (*Baz, error) {
	parsedId, err := primitives.ParseIdBytes(id[:])

	if err != nil {
		return nil, err
	}

	baz := &Baz{
		Id:    parsedId,
		Value: primitives.NewText(value),
	}

	return baz, nil
}

type Foo struct {
	Id    primitives.Id
	Value primitives.Text
	Bars  []Bar
	Bazs  []Baz
}

func NewFoo(id []byte, value string, bars []Bar, bazs []Baz) (*Foo, error) {
	parsedId, err := primitives.ParseIdBytes(id[:])

	if err != nil {
		return nil, err
	}

	foo := &Foo{
		Id:    parsedId,
		Value: primitives.NewText(value),
		Bars:  bars,
		Bazs:  bazs,
	}

	return foo, nil
}

func GenerateFoo(value string) *Foo {
	return &Foo{
		Id:    primitives.GenerateId(),
		Value: primitives.NewText(value),
		Bars:  []Bar{},
		Bazs:  []Baz{},
	}
}

func (f *Foo) AddBarBaz(num uint) {
	newBars := make([]Bar, num)
	newBazs := make([]Baz, num)

	for i := range newBars {
		newBars[i] = Bar{Id: primitives.GenerateId(), Value: primitives.NewText(fmt.Sprintf("Bar %d", i))}
		newBazs[i] = Baz{Id: primitives.GenerateId(), Value: primitives.NewText(fmt.Sprintf("Baz %d", i))}
	}

	f.Bars = append(f.Bars, newBars...)
	f.Bazs = append(f.Bazs, newBazs...)
}
