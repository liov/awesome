package main

type FooKeng struct {
}

func (*FooKeng) Foo() {

}

func NewFooKeng() FooKeng {
	return FooKeng{}
}

func main() {
	foo := NewFooKeng()
	foo.Foo()
	NewFooKeng().Foo()
}
