package main

import (
	"fmt"
	"unsafe"
)

type FooPtr1 struct {
	A string
}

type FooPtr2 struct {
	*FooPtr1
}

func main() {
	var a FooPtr1
	a.A = "hello"
	var b FooPtr2
	b.FooPtr1 = &a
	fmt.Println(unsafe.Sizeof(&a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(uintptr(*(*uint)(unsafe.Pointer(&b)))))
}
