package main

/*
#cgo CFLAGS: -mavx2
#include <immintrin.h>
#include <stdint.h>

void avx2_subtract(uint8_t *a, uint8_t *b, uint8_t *result, int length) {
    int i = 0;
    for (; i <= length - 32; i += 32) {
        __m256i va = _mm256_loadu_si256((__m256i*)&a[i]);
        __m256i vb = _mm256_loadu_si256((__m256i*)&b[i]);
        __m256i vres = _mm256_subs_epu8(va, vb);  // 使用无符号饱和减法
        _mm256_storeu_si256((__m256i*)&result[i], vres);
    }
    // 处理剩余部分
    for (; i < length; i++) {
        result[i] = a[i] > b[i] ? a[i] - b[i] : 0;
    }
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := []uint8{10, 20, 30, 40, 50, 255, 128, 64}
	b := []uint8{5, 15, 25, 35, 45, 200, 100, 60}

	result := avx2Subtract(a, b)

	fmt.Println(result)
}

func avx2Subtract(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic("Slices must have the same length")
	}

	result := make([]uint8, len(a))
	C.avx2_subtract(
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
		(*C.uint8_t)(unsafe.Pointer(&result[0])),
		C.int(len(a)),
	)

	return result
}
