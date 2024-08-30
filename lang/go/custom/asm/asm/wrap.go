package asm

import (
	"fmt"
	"unsafe"
)

func Avx2_ssd_int16(a, b []int16) int32 {
	sum := make([]int32, 8)
	avx2_ssd_int16(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&sum[0]))
	fmt.Println(a, b, sum)
	return sum[4] + sum[5] + sum[6] + sum[7]
}
