package blas

import (
	"fmt"
	"testing"
)

func Test_cblas_dasum(t *testing.T) {
	var n int64 = 3
	var x []float64 = []float64{3, 4, 5}
	var incx int64 = 1
	res := Go_cblas_dasum(n, x, incx)
	fmt.Println(res)

	if res != 12 {
		t.Error("expected 12, got ", res)
	}
}

func Test_cblas_ddot(t *testing.T) {
	res := Go_cblas_ddot(3, []float64{1, 2, 3}, 1, []float64{2, 3, 4}, 1)
	if 20 != res {
		t.Error("expected 20, got ", res)
	}
}

// func Benchmark_cblas_ddot_bench(b *testing.B) {
// 	for i := 0; i < 100; i++ {
// 		Go_cblas_ddot(3, []float64{1, 2, 3}, 1, []float64{2, 3, 4}, 1)
// 	}
// }

func Test_cblas_daxpy(t *testing.T) {
	var n int64 = 3
	var a float64 = 2.0
	var x = []float64{1, 2, 3}
	var incx int64 = 1
	var y = []float64{1, 2, 3}
	var incy int64 = 1
	Go_cblas_daxpy(n, a, x, incx, y, incy)
	if len(y) != 3 || y[0] != 3.0 || y[1] != 6.0 || y[2] != 9.0 {
		t.Error("expected [3, 6, 9], got ", y)
	}
}
