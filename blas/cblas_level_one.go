package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

//demo
int printArray(double* p, int len) {
	printf("%d, %f\n", len, p[2]);
}

*/
import "C"

//计算两个相同长度向量的叉乘。
//ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-dot#961A869B-14D9-4E4E-98FD-9CA13802C671
func Go_cblas_ddot(n int64, x []float64, incx int64, y []float64, incy int64) float64 {
	cx, lx := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	return float64(C.cblas_ddot(lx, cx, C.int(incx), cy, C.int(incy)))
}
