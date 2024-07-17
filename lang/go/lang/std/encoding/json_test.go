package json

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/log"
	"testing"
)

type Foo struct {
	A int
	B string
}

type Foo1 struct {
	Foo
	B string
}

func Test(t *testing.T) {
	var f *Foo
	fmt.Println(json.Marshal(f))
	fmt.Println(json.Unmarshal([]byte("null"), f))
	f = &Foo{}
	fmt.Println(json.Unmarshal([]byte("null"), f))
	fs := []map[string]any{}
	fmt.Println(json.Unmarshal([]byte(`[{"A":1,"B":"1"},{"A":2,"B":"2"}]`), &fs))
	fmt.Println("数组map", fs)
	fm := map[string]any{}
	fmt.Println(json.Unmarshal([]byte(`[{"A":1,"B":"1"},{"A":2,"B":"2"}]`), &fm))
	fmt.Println("单map", fm)
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

type MarshalUintStruct struct {
	U uint16 `json:"u"`
}

func TestMarshalUint(t *testing.T) {
	var u MarshalUintStruct
	fmt.Println(json.Unmarshal([]byte(`{"u":-1}`), &u))
}

type MarshalFuncStruct struct {
	Field1 int
	Field2 func()
}

// func ，chan不支持序列化，但是加上忽略标签支持; 支持反序列化
func TestMarshalFunc(t *testing.T) {
	var foo = MarshalFuncStruct{
		Field1: 10,
		Field2: func() {},
	}
	data, err := json.Marshal(&foo)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	data = []byte(`{"field1":1}`)
	err = json.Unmarshal(data, &foo)
	if err != nil {
		log.Println(err)
	}
	log.Println(&foo)
}

type MarshalChanStruct struct {
	C chan<- int `json:"c"`
}

func TestMarshalChan(t *testing.T) {
	foo := MarshalChanStruct{C: make(chan<- int, 1)}
	foo.C <- 1
	data, err := json.Marshal(&foo)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
}
