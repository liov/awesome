package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	v, ok := ctx.Value("aaa").(map[string]string)
	fmt.Println(v, ok)
}
