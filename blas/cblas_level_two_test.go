package blas

import (
	"testing"
)

func TestGo_cblas_dgbmv(t *testing.T) {
	order := CblasRowMajor
	trans := CblasNoTrans
	var m int64 = 3
	var n int64 = 2
	var kl int64 = 1
	var ku int64 = 1
	alpha := 2.0
	a := []float64{1, 2, 3, 4, 5, 6}
	var lad int64 = 1
	x := []float64{2, 3}
	var incx int64 = 1
	beta := 2.0
	y := []float64{4, 5}
	var incy int64 = 1
	// fmt.Println(y)
	// t.Error(y)
	Go_cblas_dgbmv(order, trans, m, n, kl, ku, alpha, a, lad, x, incx, beta, y, incy)
	// fmt.Println(y)
	// t.Error(y)
}
