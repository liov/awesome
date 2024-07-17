package main

import (
	"fmt"
	reflecti "github.com/hopeio/utils/reflect"
	"reflect"
	"unsafe"
)

func main() {
	var f interface{}
	f = 1
	b := (*reflecti.Eface)(unsafe.Pointer(&f))
	fmt.Println(b.Type)
	fmt.Printf("%d\n", &b.Type)
	fmt.Printf("%d\n", b.Value)
	fmt.Println(*(*[2]uintptr)(b.Value))
	v := reflect.ValueOf(&f).Elem()
	array := v.InterfaceData()
	fmt.Println(array)
	p := unsafe.Pointer(&f)
	fmt.Printf("%d\n", p)
	b1 := (*reflecti.Value)(unsafe.Pointer(&v))
	fmt.Printf("%d\n", b1.Ptr)
	fmt.Println(*(*[2]uintptr)(b1.Ptr))
}
