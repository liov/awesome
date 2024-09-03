package goat

import (
	"fmt"
	"github.com/hopeio/utils/iter"
	"slices"
	"unsafe"
)

func Avx2SsdInt16(a, b []int16) int32 {
	sum := make([]int32, 8)
	avx2_ssd_int16(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&sum[0]))
	fmt.Println(a, b, sum)
	return iter.Sum(slices.Values(sum))
}

func Check(a, b []int16) int {
	sum := 0
	for i := range a {
		sum += int(a[i]-b[i]) * int(a[i]-b[i])
	}
	return sum
}
