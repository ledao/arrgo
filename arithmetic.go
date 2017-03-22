package arrgo

import (
    "github.com/ledao/arrgo/internal"
    "math"
)

func (a *Arrf) AddC(b float64) *Arrf {
    ta := a.Copy()
    asm.AddC(b, ta.data)
    return ta
}

func(a *Arrf) Add(b *Arrf) *Arrf {
    //如果b为空或者a的轴个数比b的少，则无法计算，返回SHAPE_ERROR
    if b == nil || a.Ndims() < b.Ndims() {
        panic(SHAPE_ERROR)
    }

    ta := a.Copy()
    if b.shape[len(b.shape) -1] == ta.shape[len(ta.shape)-1] {
        asm.Add(ta.data, b.data)
        return ta
    }
    if ProductIntSlice(b.shape[0:len(b.shape)-1]) != ProductIntSlice(ta.shape[0:len(ta.shape)-1]) {
        panic(SHAPE_ERROR)
    }
    st := ta.strides[len(ta.strides)-1] * ta.shape[len(ta.shape)-1]
    for i := 0; i < len(b.data); i++ {
        asm.AddC(b.data[i], ta.data[i*st:(i+1)*st])
    }
    return ta
}

func (a *Arrf) SubC(b float64) *Arrf {
    ta := a.Copy()
    asm.SubtrC(b, ta.data)
    return ta
}

func(a *Arrf) Sub(b *Arrf) *Arrf {
    ta := a.Copy()
    if b == nil || ta.Ndims() < b.Ndims() {
        panic(SHAPE_ERROR)
    }
    if b.shape[len(b.shape) -1] == ta.shape[len(ta.shape)-1] {
        asm.Subtr(ta.data, b.data)
        return ta
    }
    if ProductIntSlice(b.shape[0:len(b.shape)-1]) != ProductIntSlice(ta.shape[0:len(ta.shape)-1]) {
        panic(SHAPE_ERROR)
    }
    st := ta.strides[len(ta.strides)-1] * ta.shape[len(ta.shape)-1]
    for i := 0; i < len(b.data); i++ {
        asm.SubtrC(b.data[i], ta.data[i*st:(i+1)*st])
    }
    return ta
}


func (a *Arrf) MulC(b float64) *Arrf {
    ta := a.Copy()
    asm.MultC(b, ta.data)
    return ta
}

func(a *Arrf) Mul(b *Arrf) *Arrf {
    ta := a.Copy()
    if b == nil || ta.Ndims() < b.Ndims() {
        panic(SHAPE_ERROR)
    }
    if b.shape[len(b.shape) -1] == ta.shape[len(ta.shape)-1] {
        asm.Mult(ta.data, b.data)
        return ta
    }
    if ProductIntSlice(b.shape[0:len(b.shape)-1]) != ProductIntSlice(ta.shape[0:len(ta.shape)-1]) {
        panic(SHAPE_ERROR)
    }
    st := ta.strides[len(ta.strides)-1] * ta.shape[len(ta.shape)-1]
    for i := 0; i < len(b.data); i++ {
        asm.MultC(b.data[i], ta.data[i*st:(i+1)*st])
    }
    return ta
}

func (a *Arrf) DivC(b float64) *Arrf {
    ta := a.Copy()
    asm.DivC(b, ta.data)
    return ta
}

func(a *Arrf) Div(b *Arrf) *Arrf {
    ta := a.Copy()
    if b == nil || ta.Ndims() < b.Ndims() {
        panic(SHAPE_ERROR)
    }
    if b.shape[len(b.shape) -1] == ta.shape[len(ta.shape)-1] {
        asm.Div(ta.data, b.data)
        return ta
    }
    if ProductIntSlice(b.shape[0:len(b.shape)-1]) != ProductIntSlice(ta.shape[0:len(ta.shape)-1]) {
        panic(SHAPE_ERROR)
    }
    st := ta.strides[len(ta.strides)-1] * ta.shape[len(ta.shape)-1]
    for i := 0; i < len(b.data); i++ {
        asm.DivC(b.data[i], ta.data[i*st:(i+1)*st])
    }
    return ta
}

func (a *Arrf) DotProd(b *Arrf) float64 {
    switch {
    case len(a.shape) == 1:
        return asm.DotProd(a.data, b.data)
    }
    panic(SHAPE_ERROR)
}


func (a *Arrf) MatProd(b *Arrf) *Arrf  {
    switch {
    case a.Ndims() ==2 && b.Ndims() ==2 && a.shape[1] == b.shape[0]:
        ret := Empty(a.shape[0], b.shape[1])
        for i := 0; i < a.shape[0]; i++ {
            for j := 0; j < a.shape[1]; j++ {
                ret.Set(a.Index(Range{i, i+1}).DotProd(b.Index(Range{0, b.shape[0]}, Range{j, j+1})), i,j)
            }
        }
        return ret
    }
    panic(SHAPE_ERROR)
}

func Abs(b *Arrf) *Arrf  {
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
    for i, v := range tb.data{
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
    for i, v := range tb.data{
        tb.data[i] = math.Cos(v)
    }
    return tb
}

func Cosh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Cosh(v)
    }
    return tb
}

func Acos(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Acos(v)
    }
    return tb
}

func Acosh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Acosh(v)
    }
    return tb
}

func Sin(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Sin(v)
    }
    return tb
}

func Sinh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Sinh(v)
    }
    return tb
}

func Asin(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Asin(v)
    }
    return tb
}

func Asinh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Asinh(v)
    }
    return tb
}

func Tan(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Tan(v)
    }
    return tb
}

func Tanh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Tanh(v)
    }
    return tb
}

func Atan(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
        tb.data[i] = math.Atan(v)
    }
    return tb
}

func Atanh(b *Arrf) *Arrf {
    var tb = b.Copy()
    for i, v := range tb.data{
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
    var t = EmptyLike(a)
    for i, v := range a.data {
        t.data[i] = math.Pow(v, b.data[i])
    }
    return t
}

func Maximum(a, b *Arrf) *Arrf {
    var t = a.Copy()
    for i, v := range t.data{
        if v < b.data[i] {
            v = b.data[i]
        }
        t.data[i] = v
    }
    return t
}

func Minimum(a, b *Arrf) *Arrf {
    var t = a.Copy()
    for i, v := range t.data{
        if v > b.data[i] {
            v = b.data[i]
        }
        t.data[i] = v
    }
    return t
}

func Mod(a, b *Arrf) *Arrf {
    var t = a.Copy()
    for i, v := range t.data{
        t.data[i] = math.Mod(v, b.data[i])
    }
    return t
}

func CopySign(a, b *Arrf) *Arrf {
    ta := Abs(a)
    sign := Sign(b)
    return ta.Mul(sign)
}