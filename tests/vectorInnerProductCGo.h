# pragma once

// #include <stdio.h>
#include <immintrin.h>


// From: https://stackoverflow.com/questions/39114159/how-do-you-load-store-from-to-an-array-of-doubles-with-gnu-c-vector-extensions
typedef double v4df __attribute__((vector_size(4 * sizeof(double))));
typedef double v2df __attribute__((vector_size(2 * sizeof(double))));


v4df load_4_doubles_intel(const double *p);
v2df load_2_doubles_intel(const double *p);


double innerProductFixed4Vector(double (*left)[4], double (*right)[4]);
