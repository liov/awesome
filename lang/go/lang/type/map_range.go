package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var n int
	for range m {
		m["d"] = 4
		n++
		fmt.Println(n)
	}
	fmt.Println(len(m))
	n = 0
	for range m {
		delete(m, "d")
		n++
		fmt.Println(n)
	}
	fmt.Println(len(m))
}
