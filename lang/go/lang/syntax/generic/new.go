package main

import "fmt"

type Foo struct {
}

func main() {
	fmt.Println(*new(*Foo))
}
