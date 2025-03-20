package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Slice struct {
	A []int `json:"a"`
}

// 底层公用
func Test(t *testing.T) {
	s := make([]Slice, 4)
	data := `{"a":[1,2]}`
	var a Slice
	json.Unmarshal([]byte(data), &a)
	s[0].A = a.A
	fmt.Println(a)
	fmt.Println(s)
	data = `{"a":[3,4]}`
	json.Unmarshal([]byte(data), &a)
	s[1].A = a.A
	fmt.Println(a)
	fmt.Println(s)
	data = `{"a":[5,6,7]}`
	json.Unmarshal([]byte(data), &a)
	s[2].A = a.A
	fmt.Println(a)
	fmt.Println(s)
	data = `{"a":[1,2]}`
	json.Unmarshal([]byte(data), &a)
	s[3].A = a.A
	fmt.Println(a)
	fmt.Println(s)
}
