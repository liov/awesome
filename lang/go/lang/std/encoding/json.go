package main

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/cherry/utils/log"
)

type Foo struct {
	A int
	B string
}

type Foo1 struct {
	Foo
	B string
}

type Foo3 struct {
	C chan<- int `json:"c"`
}

func MarshalChan() {
	foo := Foo3{C: make(chan<- int, 1)}
	foo.C <- 1
	data, err := json.Marshal(&foo)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
}

func main() {
	var f *Foo
	fmt.Println(json.Marshal(f))
	fmt.Println(json.Unmarshal([]byte("null"), f))
	f = &Foo{}
	fmt.Println(json.Unmarshal([]byte("null"), f))
	fs := []map[string]any{}
	fmt.Println(json.Unmarshal([]byte(`[{"A":1,"B":"1"},{"A":2,"B":"2"}]`), &fs))
	fmt.Println(fs)
	fm := map[string]any{}
	fmt.Println(json.Unmarshal([]byte(`[{"A":1,"B":"1"},{"A":2,"B":"2"}]`), &fm))
	fmt.Println(fm)
	fmt.Println("---------------------------------------------")
	foo1 := Foo1{
		Foo: Foo{
			A: 1,
			B: "1",
		},
		B: "2",
	}
	data, err := json.Marshal(foo1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
