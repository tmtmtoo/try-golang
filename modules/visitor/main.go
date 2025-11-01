package main

import (
	"fmt"
	"strconv"
)

type Baz struct {
	a uint64
}

func (b *Baz) Accept(visitor Visitor) error {
	return visitor.VisitBaz(b)
}

type Bar struct {
	b string
}

func (b *Bar) Accept(visitor Visitor) error {
	return visitor.VisitBar(b)
}

type Visitor interface {
	VisitBaz(baz *Baz) error
	VisitBar(bar *Bar) error
}

type Foo interface {
	Accept(visitor Visitor) error
}

type Mapper struct {
	result *string
}

func (m *Mapper) VisitBaz(b *Baz) error {
	x := strconv.FormatUint(b.a, 10)
	m.result = &x
	return nil
}

func (m *Mapper) VisitBar(b *Bar) error {
	m.result = &b.b
	return nil
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func main() {
	m := NewMapper()

	f := Foo(&Bar{b: "aaa"})

	f.Accept(m)

	fmt.Println(*m.result)
}
