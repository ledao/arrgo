package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

*/
import "C"

// res := sum(x)
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-asum
func Go_cblas_dasum(n int64, x []float64, incx int64) float64 {
	cx, _ := SliceToCArrayFloat64(x)
	return float64(C.cblas_dasum(C.int(n), cx, C.int(incx)))
}

// res := dot(x, y)
//计算两个相同长度向量的叉乘。
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-dot
func Go_cblas_ddot(n int64, x []float64, incx int64, y []float64, incy int64) float64 {
	cx, lx := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	return float64(C.cblas_ddot(lx, cx, C.int(incx), cy, C.int(incy)))
}

//y := a*x + y
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-axpy
func Go_cblas_daxpy(n int64, a float64, x []float64, incx int64, y []float64, incy int64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_daxpy(C.int(n), C.double(a), cx, C.int(incx), cy, C.int(incy))
}

// y := x
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-copy
func Go_cblas_dcopy(n int64, x []float64, incx int64, y []float64, incy int64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_dcopy(C.int(n), cx, C.int(incx), cy, C.int(incy))
}

// res = ||x||
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-nrm2
func Go_cblas_dnrm2(n int64, x []float64, incx int64) float64 {
	cx, _ := SliceToCArrayFloat64(x)
	return float64(C.cblas_dnrm2(C.int(n), cx, C.int(incx)))
}

// xi = c*xi + s*yi, yi = c*yi - s*xi
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-rot
func Go_cblas_drot(n int64, x []float64, incx int64, y []float64, incy int64, c float64, s float64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_drot(C.int(n), cx, C.int(incx), cy, C.int(incy), C.double(c), C.double(s))
}

// x = a*x
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-scal
func Go_cblas_dscal(n int64, a float64, x []float64, incx int64) {
	cx, _ := SliceToCArrayFloat64(x)
	C.cblas_dscal(C.int(n), C.double(a), cx, C.int(incx))
}

// x, y := y, x
// 交换x和y的值。
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-swap
func Go_cblas_dswap(n int64, x []float64, incx int64, y []float64, incy int64) {
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_dswap(C.int(n), cx, C.int(incx), cy, C.int(incy))
}

// res := argmax_i(abs(x[i]))
// 返回最大绝对值的位置
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-i-amax
func Go_cblas_idamax(n int64, x []float64, incx int64) int64 {
	cx, _ := SliceToCArrayFloat64(x)
	return int64(C.cblas_idamax(C.int(n), cx, C.int(incx)))
}

// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-i-amin
func Go_cblas_idamin(n int64, x []float64, incx int64) int64 {
	cx, _ := SliceToCArrayFloat64(x)
	return int64(C.cblas_idamin(C.int(n), cx, C.int(incx)))
}
