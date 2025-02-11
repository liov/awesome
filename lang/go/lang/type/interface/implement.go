package main

import (
	"fmt"
	"reflect"
)

type I interface {
	Foo()
}

var Itype = reflect.TypeOf((*I)(nil)).Elem()

type A struct {
}

func (A) Foo() {

}

type B struct {
}

func (*B) Foo() {

}

type C struct {
	B
}

func main() {
	var a A
	var b B
	var c C
	fmt.Println(reflect.TypeOf(a).Implements(Itype))
	fmt.Println(reflect.TypeOf(&a).Implements(Itype))
	fmt.Println(reflect.TypeOf(b).Implements(Itype))
	fmt.Println(reflect.TypeOf(&b).Implements(Itype))
	fmt.Println(reflect.TypeOf(c).Implements(Itype))
	fmt.Println(reflect.TypeOf(&c).Implements(Itype))
	fmt.Println("can addr:", reflect.ValueOf(c).Field(0).CanAddr())
	fmt.Println("can addr:", reflect.ValueOf(&c).Elem().Field(0).CanAddr())
	fmt.Println(reflect.ValueOf(&c).Elem().Field(0).Addr().Type().Implements(Itype))
	fmt.Println(reflect.ValueOf(&b).Elem().Addr().Type().Implements(Itype))
}
