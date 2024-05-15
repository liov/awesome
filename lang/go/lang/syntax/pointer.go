package main

import "fmt"

type Int int

func (i *Int) Set() {
	*i = 2
}

type Foo struct {
	A *Int
}

func main() {
	var a int
	ap := &a
	*ap = 1
	fmt.Println(a)
	var c Int
	c.Set()
	fmt.Println(c)
	var d Foo
	d.A.Set()
	fmt.Println(d)
	var b *int
	*b = 1
	fmt.Println(b)
}
