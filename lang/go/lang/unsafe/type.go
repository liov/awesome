package main

import (
	"fmt"
	reflecti "github.com/hopeio/utils/reflect"
	"unsafe"
)

type Foo struct{}

func main() {
	var f interface{}
	f = 1
	b := (*reflecti.Eface)(unsafe.Pointer(&f))
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
	f = Foo{}
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
}
