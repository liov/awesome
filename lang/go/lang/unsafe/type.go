package main

import (
	"fmt"
	reflectx "github.com/hopeio/gox/reflect"
	"unsafe"
)

type Foo struct{}

func main() {
	var f interface{}
	f = 1
	b := (*reflectx.Eface)(unsafe.Pointer(&f))
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
	f = Foo{}
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
}
