package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/hopeio/gox/math/bits"
	reflectx "github.com/hopeio/gox/reflect"
)

func TestFloat(t *testing.T) {
	var f float64 = 2.9999
	fmt.Println(uint64(f))
	ptr := (*uint64)(unsafe.Pointer(&f))
	fmt.Println(ptr)
	*ptr++
	fmt.Println(*(*float64)(unsafe.Pointer(ptr)))
}

func TestFloatToInt(t *testing.T) {
	var a int64 = 32
	fmt.Println(transform(a))
	bits.ViewBin(transform(1.6e-322))
	var b int32 = 32
	fmt.Println(transform(b))
	bits.ViewBin(transform(float32(4.5e-44)))
	fmt.Println(transform(int64(b)))
}

func transform(f interface{}) interface{} {
	p := (*reflectx.Eface)(unsafe.Pointer(&f)).Value
	switch f.(type) {
	case float32:
		return *(*int32)(p)
	case float64:
		return *(*int64)(p)
	case int32:
		return *(*float32)(p)
	case int64:
		return *(*float64)(p)
	}
	panic("类型不匹配")
}

type Foo struct{}

func TestType(t *testing.T) {
	var f interface{}
	f = 1
	b := (*reflectx.Eface)(unsafe.Pointer(&f))
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
	f = Foo{}
	fmt.Println(b.Type)
	fmt.Println(b.Type.Kind())
}

func TestPointer(t *testing.T) {
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
