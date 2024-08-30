package main

/*
#cgo CFLAGS: -mavx2 -std=c11
#include <immintrin.h>
#include <stdint.h>
#include <stddef.h>

uint64_t avx2_ssd(const uint8_t *a, const uint8_t *b, size_t length) {
    size_t i = 0;
    __m256i sum = _mm256_setzero_si256(); // 初始化累加器

     // 每次处理 32 个字节
    for (; i + 31 < length; i += 32) {
        __m256i va = _mm256_loadu_si256((const __m256i*)&a[i]);
        __m256i vb = _mm256_loadu_si256((const __m256i*)&b[i]);

        // 将 8 位无符号整数转换为 16 位有符号整数
        __m256i va_16 = _mm256_cvtepu8_epi16(_mm_loadu_si128((const __m128i*)&a[i]));
        __m256i vb_16 = _mm256_cvtepu8_epi16(_mm_loadu_si128((const __m128i*)&b[i]));

        // 计算差值
        __m256i vdiff = _mm256_sub_epi16(va_16, vb_16);

         __m256i vsquare = _mm256_madd_epi16(vdiff, vdiff); // 计算差值的平方

        // 使用 _mm256_add_epi64 累加差的平方和
        sum = _mm256_add_epi64(sum, _mm256_unpacklo_epi32(vsquare, _mm256_setzero_si256()));
        sum = _mm256_add_epi64(sum, _mm256_unpackhi_epi32(vsquare, _mm256_setzero_si256()));
    }

    // 累加剩余元素
    uint64_t total_sum = 0;
    for (; i < length; i++) {
        int16_t diff = (int16_t)a[i] - (int16_t)b[i];
        total_sum += (uint64_t)(diff * diff);
    }

    // 将 SIMD 累加器中的值累加到总和
    uint64_t sums[4] __attribute__((aligned(32)));
    _mm256_store_si256((__m256i*)sums, sum);
    total_sum += sums[0] + sums[1] + sums[2] + sums[3];

    return total_sum;
}

*/
import "C"
import (
	"fmt"
	"math"
	"math/rand/v2"
	"unsafe"
)

func avx2SSD(a, b []uint8) uint64 {
	if len(a) != len(b) {
		panic("Slices must have the same length")
	}
	return uint64(C.avx2_ssd((*C.uint8_t)(unsafe.Pointer(&a[0])), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.size_t(len(a))))
}

func main() {
	// 测试数据
	a, b := make([]uint8, 32), make([]uint8, 32)
	for i := 0; i < len(a); i++ {
		a[i], b[i] = uint8(rand.N(math.MaxUint8)), uint8(rand.N(math.MaxUint8))
	}
	result := avx2SSD(a, b)

	fmt.Printf("SSD result: %d\n", result)
}
