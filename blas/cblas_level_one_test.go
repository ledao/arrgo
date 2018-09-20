package blas

import "testing"

func Test_cblas_ddot(t *testing.T) {
	res := Go_cblas_ddot(3, []float64{1, 2, 3}, 1, []float64{2, 3, 4}, 1)
	if 20 != res {
		t.Error("expected 20, got ", res)
	}
}

func Benchmark_cblas_ddot_bench(b *testing.B) {
	for i := 0; i < 100; i++ {
		Go_cblas_ddot(3, []float64{1, 2, 3}, 1, []float64{2, 3, 4}, 1)
	}
}
