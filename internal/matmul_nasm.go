//+build !amd64 noasm appengine

package asm

func DotProd(a, b []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i] * b[i]
	}
	return ret
}
