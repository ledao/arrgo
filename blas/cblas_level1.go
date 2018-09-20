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

import (
	"unsafe"
)

//一维Float64 Slice转换为C的一维数组，返回数组头指针和数组长度。
//共享内存地址。
func SliceToCArrayFloat64(slice []float64) (*C.double, C.int) {
	return (*C.double)(unsafe.Pointer(&slice[0])), C.int(len(slice))
}

func Go_cblas_ddot(n int64, x []float64, incx int64, y []float64, incy int64) float64 {
	cx, lx := SliceToCArrayFloat64(x)
	C.printArray(cx, lx)
	return 0.0
}
