package arrgo

import (
	"sort"

	"github.com/ledao/arrgo/internal"
)

func (a *Arrf) Sum(axis ...int) *Arrf {
	if len(axis) == 0 || len(axis) >= a.Ndims() {
		tot := float64(0)
		for _, v := range a.data {
			tot += v
		}
		return Full(tot, 1)
	}

	sort.IntSlice(axis).Sort()
	restAxis := make([]int, len(a.shape)-len(axis))
	ta := a.Copy()
axisR:
	for i, t := 0, 0; i < len(ta.shape); i++ {
		for _, w := range axis {
			if i == w {
				continue axisR
			}
		}
		restAxis[t] = ta.shape[i]
		t++
	}

	ln := ta.strides[0]
	for k := 0; k < len(axis); k++ {
		if ta.shape[axis[k]] == 1 {
			continue
		}
		axisShape, axisSt, axis1St := ta.shape[axis[k]], ta.strides[axis[k]], ta.strides[axis[k]+1]
		if axis1St == 1 {
			asm.Hadd(uint64(axisSt), ta.data)
			ln /= axisShape
			ta.data = ta.data[:ln]
			continue
		}

		t := ta.data[0*axis1St : 1*axis1St]
		for i := 1; i < axisShape; i++ {
			asm.Vadd(t, ta.data[i*axis1St:(i+1)*axis1St])
		}
		ln /= axisShape
		ta.data = ta.data[:ln]
	}
	ta.shape = restAxis

	tmp := 1
	for i := len(restAxis); i > 0; i-- {
		ta.strides[i] = tmp
		tmp *= restAxis[i-1]
	}
	ta.strides[0] = tmp
	//ta.data = ta.data[:tmp]
	ta.strides = ta.strides[:len(restAxis)+1]
	return ta
}

func Sum(a *Arrf, axis ...int) *Arrf {
	return a.Sum(axis...)
}

func (a *Arrf) Mean(axis ...int) *Arrf {
	if len(axis) == 0 || len(axis) >= a.Ndims() {
		tot := float64(0)
		for _, v := range a.data {
			tot += v
		}
		return Full(tot/float64(a.strides[0]), 1)
	}

	sort.IntSlice(axis).Sort()
	selectShape := make([]int, len(axis))
	for i := range selectShape {
		selectShape[i] = a.shape[axis[i]]
	}
	N := ProductIntSlice(selectShape)

	ta := a.Sum(axis...)

	return ta.DivC(float64(N))
}

func Mean(a *Arrf, axis ...int) *Arrf {
	return a.Mean(axis...)
}

func (a *Arrf) Var(axis ...int) *Arrf {
	a2 := a.Mul(a).Sum(axis...)
	m := a.Mean(axis...)
	var N int
	if len(axis) == 0 || len(axis) >= a.Ndims() {
		N = ProductIntSlice(a.shape)
	} else {
		selectShape := make([]int, len(axis))
		for i := range selectShape {
			selectShape[i] = a.shape[axis[i]]
		}
		N = ProductIntSlice(selectShape)
	}

	m2 := m.Mul(m).MulC(float64(N))
	a_m_2 := a.Sum(axis...).Mul(m).MulC(2)
	return a2.Sub(a_m_2).Add(m2).DivC(float64(N))
}

func Var(a *Arrf, axis ...int) *Arrf {
	return a.Var(axis...)
}

func (a *Arrf) Std(axis ...int) *Arrf {
	return Sqrt(a.Var(axis...))
}

func Std(a *Arrf, axis ...int) *Arrf {
	return a.Std(axis...)
}
