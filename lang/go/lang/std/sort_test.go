package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
