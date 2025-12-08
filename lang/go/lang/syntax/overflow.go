package main

import (
	"fmt"
	"math"

	constraintsx "github.com/hopeio/gox/types/constraints"
)

type Foo struct {
	A, B int
}

func main() {
	fmt.Println(math.MaxInt)
	fmt.Println(math.MinInt)
	fmt.Println(f(math.MaxInt))
	fmt.Println(f(math.MinInt))
	fmt.Println(f(math.MinInt + 1))
	fmt.Println(ValueFlip(111))
}

func f(i int) int {
	return -i
}

func ValueFlip[T constraintsx.Number](i T) T {
	return -i
}
