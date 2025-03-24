package main

import (
	"fmt"
	"reflect"
)

type A2 struct {
	A int
}

func main() {
	var a = A2{1}
	var b = 2
	reflect.ValueOf(&a).Elem().Field(0).Set(reflect.ValueOf(b))
	fmt.Println(a)
	b = 3
	reflect.ValueOf(&a).Elem().Field(0).Set(reflect.ValueOf(&b).Elem())
	fmt.Println(a)
}
