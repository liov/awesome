package main

import (
	"fmt"
	"unsafe"
)

type Func func()
type Foo1 struct {
	arr []Func
	a   int
}

type Foo2 struct {
	arr []struct{}
	a   int
}

func main() {
	foo1 := Foo1{
		arr: []Func{
			func() {
				println("foo1")
			},
		},
		a: 1,
	}
	foo2 := Foo2{
		arr: []struct{}{
			{},
		},
		a: 1,
	}
	fmt.Println(unsafe.Sizeof(foo1))
	fmt.Println(unsafe.Sizeof(foo2))
}
