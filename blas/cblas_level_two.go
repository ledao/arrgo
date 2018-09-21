package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

*/
import "C"

// y := alpha*A*x + beta*y, or y := alpha*A'*x + beta*y,
// ref: https://software.intel.com/en-us/mkl-developer-reference-c-cblas-gbmv
func Go_cblas_dgbmv(order CBLAS_ORDER, trans CBLAS_TRANSPOSE, m int64, n int64, kl int64, ku int64, alpha float64, a []float64, lad int64, x []float64, incx int64, beta float64, y []float64, incy int64) {
	ca, _ := SliceToCArrayFloat64(a)
	cx, _ := SliceToCArrayFloat64(x)
	cy, _ := SliceToCArrayFloat64(y)
	C.cblas_dgbmv(C.int(order), C.int(trans), C.int(m), C.int(n), C.int(kl), C.int(ku), C.double(alpha), ca, C.int(lad), cx, C.int(incx), cy, C.int(incy))
}
