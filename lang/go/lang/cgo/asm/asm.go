package asm

import "unsafe"

func SyscallWrite_Darwin(fd int, msg string) int

func AsmCallCAdd(cfun uintptr, a, b int64) int64

//go:noescape
func _mm512_mul_to(a, b, c, n unsafe.Pointer)

func Mm512Mul(a, b, c []float32) {
	_mm512_mul_to(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&c[0]), unsafe.Pointer(uintptr(len(a))))
}
