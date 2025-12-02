package preformance

import (
	"testing"

	"github.com/hopeio/gox/cmp"
)

type Foo struct {
	A int
}

func (f *Foo) Compare(f2 *Foo) int {
	return f.A - f2.A
}

type Foo2 struct {
	Foo
}

func Compare[T cmp.Comparable[T]](a, b T) int {
	return a.Compare(b)
}

func TestCompare(t *testing.T) {
	a := Foo{1}
	b := Foo{2}
	t.Log(Compare(&a, &b))
}
