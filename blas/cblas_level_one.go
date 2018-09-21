package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

*/
import "C"

// res := sum(x)
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-asum#AE455054-52B2-460A-8887-12192E01D899
func Go_cblas_dasum(n int64, x []float64, incx int64) float64 {
	cx, _ := SliceToCArrayFloat64(x)
	return float64(C.cblas_dasum(C.int(n), cx, C.int(incx)))
}

// res := dot(x, y)
//计算两个相同长度向量的叉乘。
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-dot#961A869B-14D9-4E4E-98FD-9CA13802C671
func Go_cblas_ddot(n int64, x []float64, incx int64, y []float64, incy int64) float64 {
	cx, lx := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	return float64(C.cblas_ddot(lx, cx, C.int(incx), cy, C.int(incy)))
}

//y := a*x + y
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-axpy#C8CBB256-EAB7-4629-80FF-14029038E6B7
func Go_cblas_daxpy(n int64, a float64, x []float64, incx int64, y []float64, incy int64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_daxpy(C.int(n), C.double(a), cx, C.int(incx), cy, C.int(incy))
}

// y := x
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-copy#7E841640-7C31-4CD1-BC44-34F7D2343215
func Go_cblas_dcopy(n int64, x []float64, incx int64, y []float64, incy int64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_dcopy(C.int(n), cx, C.int(incx), cy, C.int(incy))
}

// res = ||x||
// https://software.intel.com/en-us/mkl-developer-reference-c-cblas-nrm2#ACA1A6DC-60A0-4689-BECC-2E0F4662C093
func Go_cblas_dnrm2(n int64, x []float64, incx int64) float64 {
	cx, _ := SliceToCArrayFloat64(x)
	return float64(C.cblas_dnrm2(C.int(n), cx, C.int(incx)))
}
