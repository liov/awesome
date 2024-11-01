package main

import "fmt"

func main() {
	var a = []byte{1, 2, 3}
	var b []byte
	ptr(a, &b)
	fmt.Println(b)
}

func ptr(a []byte, b any) {
	if v := b.(*[]byte); v != nil {
		*v = a[0:3]
	}
}
