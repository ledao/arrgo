package blas

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -lopenblas

#include <stdio.h>
#include "cblas.h"

*/
import "C"
import "unsafe"

//一维Float64 Slice转换为C的一维数组，返回数组头指针和数组长度。
//共享内存地址。
func SliceToCArrayFloat64(slice []float64) (*C.double, C.blasint) {
	return (*C.double)(unsafe.Pointer(&slice[0])), C.blasint(len(slice))
}
