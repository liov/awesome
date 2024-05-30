package main

import "fmt"

func main() {
	s := "☔☕♈♉♊"
	for i, n := range s {
		fmt.Println(i, n)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
