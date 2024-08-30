#include <immintrin.h>
#include <stdint.h>
#include <stddef.h>

void avx2_ssd_int16(const int16_t *a, const int16_t *b,const int32_t *sums) {

    __m256i sum = _mm256_setzero_si256(); // 初始化累加器
    __m256i va = _mm256_loadu_si256((const __m256i*)a); // 加载 16 个 int16
    __m256i vb = _mm256_loadu_si256((const __m256i*)b); // 加载 16 个 int16
    __m256i vdiff = _mm256_sub_epi16(va, vb); // 计算差值

     // 计算平方
     __m256i vsquare = _mm256_madd_epi16(vdiff, vdiff); // 计算差值的平方

    // 使用 _mm256_add_epi64 累加差的平方和
    _mm256_store_si256((__m256i*)&sums, vsquare);

}
