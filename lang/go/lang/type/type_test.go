package main

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type S1 uint32

func main() {
	var foo S1
	t := reflect.TypeOf(&foo).Elem()
	println(t.Kind())
	v := reflect.ValueOf(&foo).Elem()
	println(v.Kind())
	Tpy(foo)

	var q = Queue{}
	fmt.Printf("%p\n", &q)
	q.Point()
}

func Tpy(v interface{}) {
	switch v.(type) {
	case uint32:
		println("uint32")
	case S1:
		println("Foo")
	}
}

// 两种结构体占用的内存大小是一样的，适当的时候用适当的定义方式，当要组成新的数据结构的时候一般应该用包含的方式，
// 可以避免强转及实现指针方法时来回复制结构体
type Stack struct {
	v []interface{}
}

type Queue []interface{}

func (receiver Queue) Point() {
	fmt.Printf("%p\n", &receiver)
}

type List struct {
	list.List
}

type Func func()
type Foo1 struct {
	arr []Func
	a   int
}

type Foo2 struct {
	arr []struct{}
	a   int
}

func TestSizeof(t *testing.T) {
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

type FooPtr1 struct {
	A string
}

type FooPtr2 struct {
	*FooPtr1
}

func TestPtr(t *testing.T) {
	var a FooPtr1
	a.A = "hello"
	var b FooPtr2
	b.FooPtr1 = &a
	fmt.Println(unsafe.Sizeof(&a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(uintptr(*(*uint)(unsafe.Pointer(&b)))))
}
