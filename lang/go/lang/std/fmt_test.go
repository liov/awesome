package main

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
}
