#include <immintrin.h>
#include <stdint.h>

void Avx2SsdInt16(const int16_t *a, const int16_t *b,const int32_t *sums) {
    __m256i va = _mm256_loadu_si256((const __m256i*)a);
    __m256i vb = _mm256_loadu_si256((const __m256i*)b);
    __m256i vdiff = _mm256_sub_epi16(va, vb);

     __m256i vsquare = _mm256_madd_epi16(vdiff, vdiff);
    _mm256_storeu_si256((__m256i*)sums, vsquare);
}

