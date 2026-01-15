package _interface

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestAssert(t *testing.T) {
	var foo interface{}
	foo = 0
	if i, ok := foo.(int); ok {
		fmt.Println(i)
	}
	foo = "0"
	if i, ok := foo.(string); ok {
		fmt.Println(i)
	}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) err {
	return &errorString{text}
}

type err interface {
	Error() string
}

func TestDereference(t *testing.T) {
	var a *error
	var b *err
	c := errors.New("error")
	d := New("error")
	a = &c
	b = &d
	/*	e := &errorString{"error"}
		b = &e*/
	t.Log(*a)
	t.Log(*b)
}

type ABer1 interface {
	A()
	B()
}

type ABer2 interface {
	A()
	B()
}

type AB1 ABer1

type AB2 = ABer2

type E struct{}

type F = E

func (e *E) A() {
	fmt.Println("A")
}

func (e *E) B() {
	fmt.Println("B")
}

func TestDuck(t *testing.T) {
	var foo ABer1
	foo = &F{}
	var bar ABer2
	bar = foo
	var c AB1
	c = bar
	var d AB2
	d = c
	d.A()
	d.B()
}

type Fooer interface {
	Foo()
}

var Itype = reflect.TypeOf((*Fooer)(nil)).Elem()

type E2 struct {
}

func (E2) Foo() {

}

type E3 struct {
}

func (*E3) Foo() {

}

type CB struct {
	B
}

func TestImplement(t *testing.T) {
	var a E2
	var b E3
	var c CB
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

type Ptr interface {
	Ptr()
}

type InterPtr struct {
	Ptr *Ptr
}

func TestPtr(t *testing.T) {
	var p InterPtr
	fmt.Println(p)
}

type Adder interface {
	Add(inter Adder) Adder
}

type AdderSlice []Adder

func (f *AdderSlice) String() string {
	return "切片指针"
}

func (f AdderSlice) Add(inter Adder) Adder {
	for i := range f {
		f[i].Add(inter)
	}
	return f
}

type I int

func (i I) Add(j Adder) Adder {
	return j.(I) + i
}

func TestAdd(t *testing.T) {
	var i, j I = 1, 2
	fmt.Print(i.Add(j))
}

type fooer interface {
	foo()
}

type A struct {
	fooer
	i int
}

type B struct {
	fooer
	s string
}

type E6 struct{}

func (c *E6) foo() {
	fmt.Println("C")
}

type D struct {
	E6
	i int
}

func TestMehdodUp(t *testing.T) {
	var f fooer
	f = A{}
	fmt.Println(f)
	f = &D{}
	fmt.Println(f)
	a := A{}
	b := B{}
	a.fooer = b //这里是值拷贝，相当于一个新的A{}，用指针会stack overflow
	b.fooer = a //因为是值拷贝，所以这行并不影响变量a
	data, _ := json.Marshal(a)
	fmt.Println(string(data))
	c := A{}
	c.fooer = c
	fmt.Println(c)
	v1 := A{fooer: &E6{}}
	v1.foo()
}

type E4 interface{}

type Func func()
type FF func(Func)

type FI fooer

type BI struct {
	fooer
	Func
}

func (b *BI) foo() {
	fmt.Println("BI")
}

type BT struct{}

func (b *BT) foo() {
	fmt.Println("BT")
}

func TestStruct(t *testing.T) {
	a := BI{fooer: &BT{}, Func: (&BT{}).foo}
	a.foo()
	a.fooer.foo()
	b := BI{}
	b.fooer = &b
	b.foo()
}

/*
虽然 interface 看起来像指针类型，但它不是。interface 类型的变量只有在类型和值均为 nil 时才为 nil

如果你的 interface 变量的值是跟随其他变量变化的（雾），与 nil 比较相等时小心：
*/
func TestNil(t *testing.T) {
	if getIfac() != nil {
		fmt.Println("不为nil")
	} else {
		fmt.Println("为nil")
	}
}

type E5 struct{}

type barer interface {
	bar()
}

func (foo *E5) bar() {}

func getIfac() barer {
	foo := new(E5)
	if foo == nil {
		fmt.Println("不为nil")
	} else {
		fmt.Println("为nil")
	}
	return foo
}

type test1 struct {
	V int
}

type test2 struct {
	V int
}

// 如果接收器不是指针，则ifa接口可以是指针，也可以是对象，否则只能是指针
func (t *test1) foo() {
	t.V = 2
	fmt.Println("fooooo:", t.V)
}

func build(t *test2) fooer {
	return &test1{V: t.V}
}

func GetFoo(i fooer) {
	i.foo()
}

func TestInterface(t *testing.T) {
	var aa fooer

	aa = build(&test2{V: 1})
	fmt.Println(&aa)
	GetFoo(aa)
	fmt.Println(aa)

	var x interface{}
	b := 0
	c := "?"
	x = b
	x = c
	fmt.Println(x)

	//interfacePtr(&aa) // error
	interfacePtr(&x)
}

func interfacePtr(*any) {

}

type PrintA interface {
	GetString(b byte) string
}
type PrintB interface {
	GetString(b byte) string
}

type GetPrintA interface {
	GetPrint() PrintA
}

type GetPrintB interface {
	GetPrint() PrintB
}

type PrintC = interface {
	GetString(b byte) string
}

type GetPrintC interface {
	GetPrint() PrintC
}

type P struct {
}

func (p *P) GetString(b byte) string {
	return "Print"
}

type G struct {
}

func (g *G) GetPrint() PrintC {
	return &P{}
}

var _ PrintA = &P{}

var _ GetPrintC = &G{}

//var _ PrintA = &G{}.GetPrint()
/*
可以这么写
var b GetPrintC = &G{}
var a PrintA =b.GetPrint()
但却不能这么写
var a PrintA = &G{}.GetPrint()
cannot use &G literal.GetPrint() (type *interface { GetString(byte) string }) as type PrintA in assignment:
	*interface { GetString(byte) string } is pointer to interface, not interface
我猜
其实应该这么写
var a PrintA = (&G{}).GetPrint()
取地址运算级别最低...
*/
func TestType(t *testing.T) {
	var b GetPrintC = &G{}
	var a PrintA = b.GetPrint()
	var c PrintA = (&G{}).GetPrint()
	fmt.Println(a.GetString('b'))
	fmt.Println(c.GetString('b'))
}
