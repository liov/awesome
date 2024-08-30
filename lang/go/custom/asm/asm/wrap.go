package asm

import "unsafe"

func Avx2_ssd_int16(a, b []int16, l int) int32 {
	var result int32
	avx2_ssd_int16(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&result), unsafe.Pointer(uintptr(len(a))))
	return result
}
