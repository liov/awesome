package main

/*
#cgo CFLAGS: -mavx2 -std=c11

#include <immintrin.h>
#include <stdint.h>


void avx2_ssd_int16(const int16_t *a, const int16_t *b,const int32_t *sums) {

    __m256i va = _mm256_loadu_si256((const __m256i*)a); // 加载 16 个 int16
    __m256i vb = _mm256_loadu_si256((const __m256i*)b); // 加载 16 个 int16
    __m256i vdiff = _mm256_sub_epi16(va, vb); // 计算差值

      // 使用 _mm256_add_epi64 累加差的平方和
     __m256i vsquare = _mm256_madd_epi16(vdiff, vdiff); // 计算差值的平方


    _mm256_storeu_si256((__m256i*)sums, vsquare);

}

*/
import "C"
import (
	"fmt"
	"math"
	"math/rand/v2"

	slicesx "github.com/hopeio/gox/slices"

	"test/custom/asm/asm"
	"test/custom/asm/goat"
	"unsafe"
)

func main() {

	a, b := make([]int16, 16), make([]int16, 16)
	for i := 0; i < len(a); i++ {
		a[i], b[i] = int16(rand.N(math.MaxUint8)), int16(rand.N(math.MaxUint8))
	}

	// 调用汇编实现的 AVX2 函数
	result := asm.Avx2SsdInt16(a, b)
	// 打印部分结果
	fmt.Println("Result (first 10 values):", result)
	sum := make([]int32, 8)
	C.avx2_ssd_int16((*C.int16_t)(unsafe.Pointer(&a[0])), (*C.int16_t)(unsafe.Pointer(&b[0])), (*C.int32_t)(unsafe.Pointer(&sum[0])))
	fmt.Println("Result (first 10 values):", slicesx.Sum(sum), goat.Check(a, b))
}
