package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [5]int
	var b []int
	fmt.Println(reflect.TypeOf(a).ConvertibleTo(reflect.TypeOf(b)))
	fmt.Println(reflect.TypeOf(b).ConvertibleTo(reflect.TypeOf(a)))
}
