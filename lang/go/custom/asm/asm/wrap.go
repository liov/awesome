package asm

import (
	"unsafe"
)

func Avx2SsdInt16(a, b []int16) int32 {
	sum := make([]int32, 8)
	_Avx2SsdInt16(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&sum[0]))
	return sum[0] + sum[1] + sum[2] + sum[3] + sum[4] + sum[5] + sum[6] + sum[7]
}

func Check(a, b []int16) int {
	sum := 0
	for i := range a {
		sum += int(a[i]-b[i]) * int(a[i]-b[i])
	}
	return sum
}
