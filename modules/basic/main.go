package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	list := []uint32{1, 2, 3, 4, 5}

	sum := lo.Reduce(list, func(acc uint32, item uint32, index int) uint32 {
		return acc + item
	}, 0)

	fmt.Printf("sum: %d\n", sum)
}
