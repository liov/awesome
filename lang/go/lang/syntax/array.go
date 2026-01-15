package main

import "fmt"

func main() {
	var a Arr
	a.method()
	fmt.Println(a)
}

type Arr [5]int

func (a *Arr) method() {
	a[0] = 1
}
