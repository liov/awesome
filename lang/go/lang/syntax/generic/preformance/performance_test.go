package preformance

import (
	"github.com/hopeio/gox/cmp"
	"testing"
)

type Foo struct {
	A int
}

func (f *Foo) CompareKey() int {
	return f.A
}

func (f *Foo) Compare(f2 *Foo) int {
	return f.A - f2.A
}

type Foo2 struct {
	Foo
}

func (f *Foo2) CompareKey() int {
	return f.Foo.A
}

func Compare[T cmp.Comparable[T]](a, b T) int {
	return a.Compare(b)
}

var _ cmp.CompareKey[int] = &Foo2{}

func TestCompare(t *testing.T) {
	a := Foo{1}
	b := Foo{2}
	t.Log(Compare(&a, &b))
}
