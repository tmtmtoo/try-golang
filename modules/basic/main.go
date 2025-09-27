package main

import (
	"basic/example"
	"basic/nesting"
	"basic/nesting/nested"
	"basic/protecting"

	// Alias to avoid name conflict
	nested2 "basic/protecting/nested"

	// Inaccessible due to internal directory
	// "basic/protecting/internal/protected"

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

	nestingString := nesting.ExampleNestingFunction()
	fmt.Println(nestingString)

	nestedString := nested.ExampleNestedFunction()
	fmt.Println(nestedString)

	protectingString := protecting.ExampleProtectingFunction()
	fmt.Println(protectingString)

	// Inaccessible due to internal directory
	// protected := protected.ExampleProtectedFunction()
	// fmt.Println(protected)

	nested2String := nested2.ExampleNestedFunction2()
	fmt.Println(nested2String)
}
