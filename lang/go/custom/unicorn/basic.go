package main

import (
	"fmt"
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

func main() {
	mu, err := uc.NewUnicorn(uc.ARCH_X86, uc.MODE_32)
	if err != nil {
		panic(err)
	}
	// mov eax, 1234
	code := []byte{184, 210, 4, 0, 0}
	mu.MemMap(0x1000, 0x1000)
	mu.MemWrite(0x1000, code)
	if err := mu.Start(0x1000, 0x1000+uint64(len(code))); err != nil {
		panic(err)
	}
	eax, _ := mu.RegRead(uc.X86_REG_EAX)
	fmt.Printf("EAX is now: %d\n", eax)
}
