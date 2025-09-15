package main

import (
	"basic/example"
	"fmt"

	"github.com/samber/lo"
)

func exampleUseLibrary(values *[]uint32) uint32 {
	return lo.Sum(*values)
}

func main() {
	sum := exampleUseLibrary(&[]uint32{1, 2, 3, 4, 5})
	fmt.Printf("sum: %d\n", sum)

	foo := example.NewFoo("World")
	fmt.Println(foo.Greet())
}
