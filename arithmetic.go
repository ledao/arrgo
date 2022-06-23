package arrgo

import (
	"math"

	asm "github.com/ledao/arrgo/internal"
)

//多维数组和标量相加，结果为新的多维数组，不修改原数组。
func (a *Arrf) AddC(b float64) *Arrf {
	if a == nil || a.Size() == 0 {
		panic(SHAPE_ERROR)
	}
	ta := a.Copy()
	asm.AddC(b, ta.data)
	return ta
}

//两个多维数组相加，结果为新的多维数组，不修改原数组。
//加法过程中间会发生广播，对矩阵运算有极大帮助。
//fixme : by ledao 广播机制会进行额外的运算，对于简单的场景最好有判断，避免广播。
func (a *Arrf) Add(b *Arrf) *Arrf {
	if a.SameShapeTo(b) {
		var ta = a.Copy()
		asm.Add(ta.data, b.data)
		return ta
	}
	var ta, tb, err = Boardcast(a, b)
	if err != nil {
		panic(err)
	}
	return ta.Add(tb)
}

//多维数组和标量相减，结果为新的多维数组，不修改原数组。
func (a *Arrf) SubC(b float64) *Arrf {
	ta := a.Copy()
	asm.SubtrC(b, ta.data)
	return ta
}

//两个多维数组相减，结果为新的多维数组，不修改原数组。
//减法过程中间会发生广播，对矩阵运算有极大帮助。
//fixme : by ledao 广播机制会进行额外的运算，对于简单的场景最好有判断，避免广播。
func (a *Arrf) Sub(b *Arrf) *Arrf {
	if a.SameShapeTo(b) {
		var ta = a.Copy()
		asm.Subtr(ta.data, b.data)
		return ta
	}
	var ta, tb, err = Boardcast(a, b)
	if err != nil {
		panic(err)
	}
	return ta.Sub(tb)
}

func (a *Arrf) MulC(b float64) *Arrf {
	ta := a.Copy()
	asm.MultC(b, ta.data)
	return ta
}

func (a *Arrf) Mul(b *Arrf) *Arrf {
	if a.SameShapeTo(b) {
		var ta = a.Copy()
		asm.Mult(ta.data, b.data)
		return ta
	}
	var ta, tb, err = Boardcast(a, b)
	if err != nil {
		panic(err)
	}
	return ta.Mul(tb)
}

func (a *Arrf) DivC(b float64) *Arrf {
	ta := a.Copy()
	asm.DivC(b, ta.data)
	return ta
}

func (a *Arrf) Div(b *Arrf) *Arrf {
	if a.SameShapeTo(b) {
		var ta = a.Copy()
		asm.Div(ta.data, b.data)
		return ta
	}
	var ta, tb, err = Boardcast(a, b)
	if err != nil {
		panic(err)
	}
	return ta.Div(tb)
}

func (a *Arrf) DotProd(b *Arrf) float64 {
	switch {
	case a.Ndims() == 1 && b.Ndims() == 1 && a.Length() == b.Length():
		return asm.DotProd(a.data, b.data)
	}
	panic(SHAPE_ERROR)
}

func (a *Arrf) MatProd(b *Arrf) *Arrf {
	switch {
	case a.Ndims() == 2 && b.Ndims() == 2 && a.shape[1] == b.shape[0]:
		ret := Zeros(a.shape[0], b.shape[1])
		for i := 0; i < a.shape[0]; i++ {
			for j := 0; j < a.shape[1]; j++ {
				ret.Set(a.Index(Range{i, i + 1}).DotProd(b.Index(Range{0, b.shape[0]}, Range{j, j + 1})), i, j)
			}
		}
		return ret
	}
	panic(SHAPE_ERROR)
}

func Abs(b *Arrf) *Arrf {
	tb := b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Abs(v)
	}
	return tb
}

func Sqrt(b *Arrf) *Arrf {
	tb := b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Sqrt(v)
	}
	return tb
}

func Square(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Pow(v, 2)
	}
	return tb
}

func Exp(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Exp(v)
	}
	return tb
}

func Log(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Log(v)
	}
	return tb
}

func Log10(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Log10(v)
	}
	return tb
}

func Log2(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Log2(v)
	}
	return tb
}

func Log1p(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Log1p(v)
	}
	return tb
}

func Sign(b *Arrf) *Arrf {
	var tb = b.Copy()
	var sign float64 = 0
	for i, v := range tb.data {
		if v > 0 {
			sign = 1
		} else if v < 0 {
			sign = -1
		}
		tb.data[i] = sign
	}
	return tb
}

func Ceil(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Ceil(v)
	}
	return tb
}

func Floor(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Floor(v)
	}
	return tb
}

func Round(b *Arrf, places int) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = Roundf(v, places)
	}
	return tb
}

func Modf(b *Arrf) (*Arrf, *Arrf) {
	var tb = b.Copy()
	var tbFrac = b.Copy()
	for i, v := range tb.data {
		r, f := math.Modf(v)
		tb.data[i] = r
		tbFrac.data[i] = f
	}
	return tb, tbFrac
}

func IsNaN(b *Arrf) *Arrb {
	var tb = EmptyB(b.shape...)
	for i, v := range b.data {
		tb.data[i] = math.IsNaN(v)
	}
	return tb
}

func IsInf(b *Arrf) *Arrb {
	var tb = EmptyB(b.shape...)
	for i, v := range b.data {
		tb.data[i] = math.IsInf(v, 0)
	}
	return tb
}

func IsFinit(b *Arrf) *Arrb {
	var tb = EmptyB(b.shape...)
	for i, v := range b.data {
		tb.data[i] = !math.IsInf(v, 0)
	}
	return tb
}

func Cos(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Cos(v)
	}
	return tb
}

func Cosh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Cosh(v)
	}
	return tb
}

func Acos(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Acos(v)
	}
	return tb
}

func Acosh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Acosh(v)
	}
	return tb
}

func Sin(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Sin(v)
	}
	return tb
}

func Sinh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Sinh(v)
	}
	return tb
}

func Asin(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Asin(v)
	}
	return tb
}

func Asinh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Asinh(v)
	}
	return tb
}

func Tan(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Tan(v)
	}
	return tb
}

func Tanh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Tanh(v)
	}
	return tb
}

func Atan(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Atan(v)
	}
	return tb
}

func Atanh(b *Arrf) *Arrf {
	var tb = b.Copy()
	for i, v := range tb.data {
		tb.data[i] = math.Atanh(v)
	}
	return tb
}

func Add(a, b *Arrf) *Arrf {
	return a.Add(b)
}

func Sub(a, b *Arrf) *Arrf {
	return a.Sub(b)
}

func Mul(a, b *Arrf) *Arrf {
	return a.Mul(b)
}

func Div(a, b *Arrf) *Arrf {
	return a.Div(b)
}

func Pow(a, b *Arrf) *Arrf {
	var t = ZerosLike(a)
	for i, v := range a.data {
		t.data[i] = math.Pow(v, b.data[i])
	}
	return t
}

func Maximum(a, b *Arrf) *Arrf {
	var t = a.Copy()
	for i, v := range t.data {
		if v < b.data[i] {
			v = b.data[i]
		}
		t.data[i] = v
	}
	return t
}

func Minimum(a, b *Arrf) *Arrf {
	var t = a.Copy()
	for i, v := range t.data {
		if v > b.data[i] {
			v = b.data[i]
		}
		t.data[i] = v
	}
	return t
}

func Mod(a, b *Arrf) *Arrf {
	var t = a.Copy()
	for i, v := range t.data {
		t.data[i] = math.Mod(v, b.data[i])
	}
	return t
}

func CopySign(a, b *Arrf) *Arrf {
	ta := Abs(a)
	sign := Sign(b)
	return ta.Mul(sign)
}

func Boardcast(a, b *Arrf) (*Arrf, *Arrf, error) {
	if a.Ndims() < b.Ndims() {
		return nil, nil, SHAPE_ERROR
	}
	var bNewShape []int
	if a.Ndims() == b.Ndims() {
		bNewShape = b.shape
	} else {
		bNewShape = make([]int, len(a.shape))
		for i := range bNewShape {
			bNewShape[i] = 1
		}
		copy(bNewShape[len(a.shape)-len(b.shape):], b.shape)
	}

	var aChangeIndex = make([]int, 0)
	var aChangeNum = make([]int, 0)
	var bChangeIndex = make([]int, 0)
	var bChangeNum = make([]int, 0)
	for i := range bNewShape {
		if a.shape[i] == bNewShape[i] {
			continue
		} else if a.shape[i] == 1 {
			aChangeIndex = append(aChangeIndex, i)
			aChangeNum = append(aChangeNum, bNewShape[i])
		} else if bNewShape[i] == 1 {
			bChangeIndex = append(bChangeIndex, i)
			bChangeNum = append(bChangeNum, a.shape[i])
		} else {
			return nil, nil, SHAPE_ERROR
		}
	}

	var aNew, bNew *Arrf
	if len(aChangeNum) == 0 {
		aNew = a
	} else {
		var baseNum = a.Length()
		var expandTimes = ProductIntSlice(aChangeNum)
		var expandData = make([]float64, baseNum*expandTimes)
		for i := 0; i < expandTimes; i++ {
			copy(expandData[i*baseNum:(i+1)*baseNum], a.data)
		}
		var newPos = make([]int, len(aChangeIndex), len(a.shape))
		var expandShape = make([]int, len(aChangeNum), len(a.shape))
		copy(newPos, aChangeIndex)
		copy(expandShape, aChangeNum)
		for i := range a.shape {
			if !ContainsInt(aChangeIndex, i) {
				newPos = append(newPos, i)
				expandShape = append(expandShape, a.shape[i])
			}
		}
		aNew = Array(expandData, expandShape...).Transpose(newPos...)
	}

	if len(bChangeNum) == 0 {
		bNew = b
	} else {
		var baseNum = b.Length()
		var expandTimes = ProductIntSlice(bChangeNum)
		var expandData = make([]float64, baseNum*expandTimes)
		for i := 0; i < expandTimes; i++ {
			copy(expandData[i*baseNum:(i+1)*baseNum], b.data)
		}
		var newPos = make([]int, len(bChangeIndex), len(bNewShape))
		var expandShape = make([]int, len(bChangeNum), len(bNewShape))
		copy(newPos, bChangeIndex)
		copy(expandShape, bChangeNum)
		for i := range bNewShape {
			if !ContainsInt(bChangeIndex, i) {
				newPos = append(newPos, i)
				expandShape = append(expandShape, bNewShape[i])
			}
		}
		bNew = Array(expandData, expandShape...).Transpose(newPos...)
	}

	return aNew, bNew, nil
}
