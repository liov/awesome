package main

import (
	"context"
	"fmt"
	"testing"
)

func TestAssert(t *testing.T) {
	ctx := context.Background()
	v, ok := ctx.Value("aaa").(map[string]string)
	fmt.Println(v, ok)
}
