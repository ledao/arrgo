package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

*/
import "C"

func Go_cblas_ddot(n int64, x []float64, incx int64, y []float64, incy int64) float64 {

	return 0.0
}
