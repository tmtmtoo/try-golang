package foo

import "fmt"

type Bar struct {
	Id    Id
	Value Text
}

func NewBar(id []byte, value string) (*Bar, error) {
	parsedId, err := ParseIdBytes(id)

	if err != nil {
		return nil, err
	}

	bar := &Bar{
		Id:    parsedId,
		Value: Text(value),
	}

	return bar, nil
}

type Baz struct {
	Id    Id
	Value Text
}

func NewBaz(id []byte, value string) (*Baz, error) {
	parsedId, err := ParseIdBytes(id[:])

	if err != nil {
		return nil, err
	}

	baz := &Baz{
		Id:    parsedId,
		Value: Text(value),
	}

	return baz, nil
}

type Foo struct {
	Id    Id
	Value Text
	Bars  []Bar
	Bazs  []Baz
}

func NewFoo(id []byte, value string, bars []Bar, bazs []Baz) (*Foo, error) {
	parsedId, err := ParseIdBytes(id[:])

	if err != nil {
		return nil, err
	}

	foo := &Foo{
		Id:    parsedId,
		Value: Text(value),
		Bars:  bars,
		Bazs:  bazs,
	}

	return foo, nil
}

func GenerateFoo(value string) *Foo {
	return &Foo{
		Id:    GenerateId(),
		Value: Text(value),
		Bars:  []Bar{},
		Bazs:  []Baz{},
	}
}

func (f *Foo) AddBarBaz(num uint) {
	newBars := make([]Bar, num)
	newBazs := make([]Baz, num)

	for i := range newBars {
		newBars[i] = Bar{Id: GenerateId(), Value: Text(fmt.Sprintf("Bar %d", i))}
		newBazs[i] = Baz{Id: GenerateId(), Value: Text(fmt.Sprintf("Baz %d", i))}
	}

	f.Bars = append(f.Bars, newBars...)
	f.Bazs = append(f.Bazs, newBazs...)
}
