package blas

import "testing"

func Test_cblas_ddot(t *testing.T) {
	res := Go_cblas_ddot(1, []float64{1, 2, 3}, 1, []float64{2}, 1)
	if 0.0 != res {
		t.Error("expected 0.0, got ", res)
	}
}
