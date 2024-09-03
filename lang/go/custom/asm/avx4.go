package main

/*
#cgo CFLAGS: -mavx2 -std=c11
// 文件名：avx2_ssd_int16.c
#include <immintrin.h>
#include <stdint.h>
#include <stddef.h>

uint64_t avx2_ssd_int16(const int16_t *a, const int16_t *b, size_t length) {
    size_t i = 0;
    __m256i sum = _mm256_setzero_si256(); // 初始化累加器

    // 每次处理 16 个 int16（32 字节）
    for (; i + 15 < length; i += 16) {
        __m256i va = _mm256_loadu_si256((const __m256i*)&a[i]); // 加载 16 个 int16
        __m256i vb = _mm256_loadu_si256((const __m256i*)&b[i]); // 加载 16 个 int16
        __m256i vdiff = _mm256_sub_epi16(va, vb); // 计算差值
        // 计算平方
         __m256i vsquare = _mm256_madd_epi16(vdiff, vdiff); // 计算差值的平方

        // 使用 _mm256_add_epi64 累加差的平方和
        sum = _mm256_add_epi64(sum, _mm256_unpacklo_epi32(vsquare, _mm256_setzero_si256()));
        sum = _mm256_add_epi64(sum, _mm256_unpackhi_epi32(vsquare, _mm256_setzero_si256()));
    }

    // 累加剩余元素
    uint64_t total_sum = 0;
    for (; i < length; i++) {
        int32_t diff = (int32_t)a[i] - (int32_t)b[i];
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
	"test/custom/asm/goat"
	"unsafe"
)

func avx2SSDInt16(a, b []int16) uint64 {
	if len(a) != len(b) {
		panic("Slices must have the same length")
	}
	return uint64(C.avx2_ssd_int16((*C.int16_t)(unsafe.Pointer(&a[0])), (*C.int16_t)(unsafe.Pointer(&b[0])), C.size_t(len(a))))
}

func main() {
	// 测试数据
	a, b := make([]int16, 32), make([]int16, 32)
	for i := 0; i < len(a); i++ {
		a[i], b[i] = int16(rand.N(math.MaxUint8)), int16(rand.N(math.MaxUint8))
	}

	result := avx2SSDInt16(a, b)

	fmt.Printf("SSD result: %d %d\n", result, goat.Check(a, b))
}
