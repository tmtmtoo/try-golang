package protecting

import (
	"basic/protecting/internal/protected"
	"fmt"
)

func ExampleProtectingFunction() string {
	return fmt.Sprintf("example protecting function and %s", protected.ExampleProtectedFunction())
}
