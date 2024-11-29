package main

import (
	"fmt"
	"reflect"
)

type A1 struct{}

func (receiver *A1) name() {

}

func (receiver *A1) Name() {

}

func (receiver *A1) Foo() {

}

func main() {
	var a A1
	v := reflect.ValueOf(&a)
	for i := range v.NumMethod() {
		fmt.Println(i, v.Method(i).Type().String())
	}
	fmt.Println(v.MethodByName("name").String())
	t := reflect.TypeOf(&a)
	for i := range t.NumMethod() {
		fmt.Println(i, t.Method(i).Name)
	}
	fmt.Println(t.MethodByName("name"))
}
