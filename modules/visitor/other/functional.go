package other

import (
	"fmt"
	"strconv"
)

type Bar struct {
	a uint64
}

type Baz struct {
	b string
}

type Visitor struct {
	VisitBar func(bar *Bar) error
	VisitBaz func(baz *Baz) error
}

type Foo interface {
	Accept(visitor *Visitor) error
}

func (bar *Bar) Accept(visitor *Visitor) error {
	return visitor.VisitBar(bar)
}

func (baz *Baz) Accept(visitor *Visitor) error {
	return visitor.VisitBaz(baz)
}

func Hoge() {
	foo := Foo(&Baz{b: "hogehoge"})

	var res string
	foo.Accept(&Visitor{
		VisitBar: func(bar *Bar) error {
			res = strconv.FormatUint(bar.a, 10)
			return nil
		},
		VisitBaz: func(baz *Baz) error {
			res = baz.b
			return nil
		},
	})

	fmt.Println(res)
}
