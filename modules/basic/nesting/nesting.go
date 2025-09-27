package nesting

import (
	"basic/nesting/nested"
	"fmt"
)

func ExampleNestingFunction() string {
	return fmt.Sprintf("example nesting function and %s", nested.ExampleNestedFunction())
}
