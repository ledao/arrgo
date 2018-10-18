package blas

import (
	"testing"
)

func Test_cblas_dasum(t *testing.T) {
	var n int64 = 3
	var x []float64 = []float64{3, 4, 5}
	var incx int64 = 1
	res := Go_cblas_dasum(n, x, incx)

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

func Test_Go_cblas_dcopy(t *testing.T) {
	var n int64 = 3
	var x = []float64{2,3,4}
	var incx  int64 = 1
	var y = []float64{0,0,0}
	var incy int64 = 1
	Go_cblas_dcopy(n, x, incx, y, incy)
	if y[0] != 2 || y[1] != 3 || y[2] != 4 {
		t.Error("expected [2,3,4], got ", y)
	}
	x[0] = 10
	if y[0] != 2 {
		t.Error("expected 2, got ", y[0])
	}
}

func Test_Go_cblas_dnrm2(t *testing.T) {
	var n int64 = 3
	var x = []float64{2,3,4}
	var incx int64 = 1
	var norm = Go_cblas_dnrm2(n, x, incx)
	if norm != 5.385164807134504 {
		t.Error("expected 5.385164807134504, got ", norm)
	}
}

func Test_Go_cblas_drot(t *testing.T) {
	var n int64 = 3
	var x = []float64{2,3,4}
	var incx int64 = 1
	var y = []float64{2,3,4}
	var incy int64 = 1
	var c float64 = 2.0
	var s float64 = 3.0
	Go_cblas_drot(n, x, incx, y, incy, c ,s )
	if x[0] != 10 || x[1] != 15 || x[2] != 20 {
		t.Error("expected [10, 15, 20], got ", x)
	}
	if y[0] != -2 || y[1] != -3 || y[2] != -4 {
		t.Error("Expected [-2, -3, -4], got ", y)
	}
}