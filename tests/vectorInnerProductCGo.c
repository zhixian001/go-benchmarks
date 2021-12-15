#include "vectorInnerProductCGo.h"

v4df load_4_doubles_intel(const double *p) { return _mm256_loadu_pd(p); }
v2df load_2_doubles_intel(const double *p) { return _mm_loadu_pd(p); }


// v4df avx_constant() { return _mm256_setr_pd( 1.0, 2.0, 3.0, 4.0 ); }

double innerProductFixed4Vector(double (*left)[4], double (*right)[4]) {
    v2df leftTop = load_2_doubles_intel(&(*left)[0]);
    v2df leftBottom = load_2_doubles_intel(&(*left)[2]);

    v2df rightTop = load_2_doubles_intel(&(*right)[0]);
    v2df rightBottom = load_2_doubles_intel(&(*right)[2]);

    __m128d top = _mm_dp_pd(leftTop, rightTop, 0b11111111);
    __m128d bottom = _mm_dp_pd(leftBottom, rightBottom, 0b11111111);

    // printf("%f", _mm_cvtsd_f64(leftTop));
    // printf("%f", _mm_cvtsd_f64(leftBottom));

    return _mm_cvtsd_f64(top) + _mm_cvtsd_f64(bottom);
}
