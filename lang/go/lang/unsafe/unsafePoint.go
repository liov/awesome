package main

import (
	"fmt"
	reflectx "github.com/hopeio/gox/reflect"
	"reflect"
	"unsafe"
)

func main() {
	var f interface{}
	f = 1
	b := (*reflectx.Eface)(unsafe.Pointer(&f))
	fmt.Println(b.Type)
	fmt.Printf("%d\n", &b.Type)
	fmt.Printf("%d\n", b.Value)
	fmt.Println(*(*[2]uintptr)(b.Value))
	v := reflect.ValueOf(&f).Elem()
	array := v.InterfaceData()
	fmt.Println(array)
	p := unsafe.Pointer(&f)
	fmt.Printf("%d\n", p)
	b1 := (*reflectx.Value)(unsafe.Pointer(&v))
	fmt.Printf("%d\n", b1.Ptr)
	fmt.Println(*(*[2]uintptr)(b1.Ptr))
}
