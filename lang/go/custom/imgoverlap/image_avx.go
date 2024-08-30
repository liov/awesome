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

         // 计算平方
        __m256i vsquare = _mm256_mullo_epi16(vdiff, vdiff);

        // 累加平方值
        sum = _mm256_add_epi64(sum, _mm256_cvtepu32_epi64(_mm256_castsi256_si128(vsquare)));
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
	"image"
	"image/jpeg"
	"math"
	"os"
	"time"
	"unsafe"
)

func avx2SSD(a, b []uint8) uint64 {
	if len(a) != len(b) {
		panic("Slices must have the same length")
	}
	return uint64(C.avx2_ssd((*C.uint8_t)(unsafe.Pointer(&a[0])), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.size_t(len(a))))
}

// 恍然大悟，图片其实就是一维数组
func calculate(y int, img1, img2 []uint8, minOverlap, maxOverlap int) int {
	n := len(img1)
	minMean := uint64(math.MaxUint64)
	var overlap int

	for o := minOverlap; o <= maxOverlap; o++ {
		var sum uint64
		m := o * y
		subimg1 := img1[n-m:]
		subimg2 := img2[:m]
		sum = avx2SSD(subimg1, subimg2)
		if sum < minMean {
			minMean = sum
			overlap = o
			fmt.Println(overlap, sum)
		} else {
			break
		}
	}

	return overlap
}

func main() {
	now := time.Now()
	file1, err := os.Open(`D:\work\scan-panel-0826\0--light1.jpg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file2, err := os.Open(`D:\work\scan-panel-0826\1--light1.jpg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	img1, err := jpeg.Decode(file1)
	img2, err := jpeg.Decode(file2)

	minOverlap := 1200
	maXOverlap := 1300
	data1, data2 := ImageData(img1, img2, maXOverlap)
	file1.Close()
	file2.Close()
	bounds1 := img1.Bounds()
	fmt.Println(calculate(bounds1.Dy(), data1, data2, minOverlap, maXOverlap), time.Since(now))
}

func ImageData(img1, img2 image.Image, maXOverlap int) ([]uint8, []uint8) {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	data1 := make([]uint8, 0, maXOverlap*bounds1.Dy())
	// 遍历原始图像的每个像素并转换为灰度值
	for x := bounds1.Max.X - maXOverlap; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data1 = append(data1, gray)
		}
	}
	data2 := make([]uint8, 0, maXOverlap*bounds2.Dy())
	for x := bounds2.Min.X; x <= maXOverlap; x++ {
		for y := bounds2.Min.Y; y < bounds2.Max.Y; y++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data2 = append(data2, gray)
		}
	}
	return data1, data2
}
