#include <immintrin.h>
#include <stdint.h>
#include <stddef.h>

void avx2_ssd_int16(const int16_t *a, const int16_t *b, int32_t *c,size_t length) {
    size_t i = 0;
    __m256i sum = _mm256_setzero_si256(); // 初始化累加器

    // 每次处理 16 个 int16（32 字节）
    for (; i + 15 < length; i += 16) {
        __m256i va = _mm256_loadu_si256((const __m256i*)&a[i]); // 加载 16 个 int16
        __m256i vb = _mm256_loadu_si256((const __m256i*)&b[i]); // 加载 16 个 int16
        __m256i vdiff = _mm256_sub_epi16(va, vb); // 计算差值

         // 计算平方
        __m256i diff_lo = _mm256_unpacklo_epi16(vdiff, _mm256_setzero_si256());
        __m256i diff_hi = _mm256_unpackhi_epi16(vdiff, _mm256_setzero_si256());

        __m256i squared_lo = _mm256_madd_epi16(diff_lo, diff_lo);
        __m256i squared_hi = _mm256_madd_epi16(diff_hi, diff_hi);

        // 累加平方
        sum = _mm256_add_epi32(sum, squared_lo);
        sum = _mm256_add_epi32(sum, squared_hi);

    }

    // 累加剩余元素
    int32_t total_sum = 0;
    for (; i < length; i++) {
        int32_t diff = (int32_t)a[i] - (int32_t)b[i];
        total_sum += diff * diff;
    }

    // 将 SIMD 累加器中的值累加到总和
     __m128i final_sum = _mm256_castsi256_si128(sum);
    final_sum = _mm_add_epi32(final_sum, _mm_srli_si128(final_sum, 8)); // 高 64 位累加
    final_sum = _mm_add_epi32(final_sum, _mm_srli_si128(final_sum, 4)); // 高 32 位累加

    total_sum += _mm_cvtsi128_si32(final_sum);

    *c = total_sum;
}
