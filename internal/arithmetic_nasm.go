//+build !amd64 noasm appengine

package asm

var (
	Sse3Supt, AvxSupt, Avx2Supt, FmaSupt bool
)

func initasm() {
}

func AddC(c float64, d []float64) {
	for i := range d {
		d[i] += c
	}
}

func SubtrC(c float64, d []float64) {
	for i := range d {
		d[i] -= c
	}
}

func MultC(c float64, d []float64) {
	for i := range d {
		d[i] *= c
	}
}

func DivC(c float64, d []float64) {
	for i := range d {
		d[i] /= c
	}
}

func Add(a, b []float64) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] += b[j]
	}
}

func Vadd(a, b []float64) {
	for i := range a {
		a[i] += b[i]
	}
}

func Hadd(st uint64, a []float64) {
	ln := uint64(len(a))
	for k := uint64(0); k < ln/st; k++ {
		a[k] = a[k*st]
		for i := uint64(1); i < st; i++ {
			a[k] += a[k*st+i]
		}
	}
}

func Subtr(a, b []float64) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] -= b[j]
	}
}

func Mult(a, b []float64) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] *= b[j]
	}
}

func Div(a, b []float64) {
	lna, lnb := len(a), len(b)
	for i, j := 0, 0; i < lna; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		a[i] /= b[j]
	}
}

func Fma12(a float64, x, b []float64) {
	lnx, lnb := len(x), len(b)
	for i, j := 0, 0; i < lnx; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		x[i] = a*x[i] + b[j]
	}
}

func Fma21(a float64, x, b []float64) {
	lnx, lnb := len(x), len(b)
	for i, j := 0, 0; i < lnx; i, j = i+1, j+1 {
		if j >= lnb {
			j = 0
		}
		x[i] = x[i]*b[j] + a
	}
}
