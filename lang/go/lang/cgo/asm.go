package main

/*
#include <stdint.h>

int64_t myadd(int64_t a, int64_t b) {
    return a+b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"

	asmpkg "test/lang/cgo/asm"
)

func main() {

	println(asmpkg.AsmCallCAdd(
		uintptr(unsafe.Pointer(C.myadd)),
		123, 456,
	))
	asmpkg.SyscallWrite_Darwin(1, "hello syscall!\n")
	a := []float32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160}
	b := []float32{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80}
	c := make([]float32, len(a))
	asmpkg.Mm512Mul(a, b, c)

	fmt.Println(c)
}
