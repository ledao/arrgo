package arrgo


func (a *Arrf) Greater(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v > b.data[i]
    }
    return t
}

func (a *Arrf) GreaterEqual(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v >= b.data[i]
    }
    return t
}

func (a *Arrf) Less(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v < b.data[i]
    }
    return t
}

func (a *Arrf) LessEqual(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v <= b.data[i]
    }
    return t
}

func (a *Arrf) Equal(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v == b.data[i]
    }
    return t
}

func (a *Arrf) NotEqual(b *Arrf) *Arrb {
    if len(a.data) == 0 || len(b.data) == 0 {
        panic(EMPTY_ARRAY_ERROR)
    }
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v != b.data[i]
    }
    return t
}
func  Greater(a, b *Arrf) *Arrb {
    return a.Greater(b)
}

func  GreaterEqual(a,b *Arrf) *Arrb {
    return a.GreaterEqual(b)
}

func Less(a, b *Arrf) *Arrb {
    return a.Less(b)
}

func  LessEqual(a, b *Arrf) *Arrb {
    return a.LessEqual(b)
}

func Equal(a, b *Arrf) *Arrb {
    return a.Equal(b)
}

func NotEqual(a, b *Arrf) *Arrb {
    return a.NotEqual(b)
}

func (a *Arrf) Sort(axis ...int) *Arrf {
    ax := -1
    if len(axis) == 0 {
        ax = a.Ndims() - 1
    } else {
        ax = axis[0]
    }

    axisShape, axisSt, axis1St := a.shape[ax], a.strides[ax], a.strides[ax + 1]
    if axis1St == 1 {
        Hsort(axisSt, a.data)
    } else {
        Vsort(axis1St, a.data[0:axisShape * axis1St])
    }

    return a
}

func Sort(a *Arrf, axis ...int) *Arrf {
    return a.Copy().Sort(axis...)
}