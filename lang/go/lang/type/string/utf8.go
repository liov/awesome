package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱中国"
	for i, r := range s {
		println(i, r, utf8.RuneLen(r))
	}
	fmt.Println("-----------")
	s = "☔☕♈♉♊"
	for i, n := range s {
		fmt.Println(i, n)
	}
	fmt.Println("-----------")
	for i := range s {
		fmt.Println(i)
	}
	fmt.Println("-----------")
	for i := range len(s) {
		fmt.Println(i, s[i])
	}

}
