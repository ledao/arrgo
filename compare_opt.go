package arrgo


func (a *Arrf) Greater(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v > b.data[i]
    }
    return t
}

func (a *Arrf) GreaterEqual(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v >= b.data[i]
    }
    return t
}

func (a *Arrf) Less(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v < b.data[i]
    }
    return t
}

func (a *Arrf) LessEqual(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v <= b.data[i]
    }
    return t
}

func (a *Arrf) Equal(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v == b.data[i]
    }
    return t
}

func (a *Arrf) NotEqual(b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v != b.data[i]
    }
    return t
}
func  Greater(a, b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v > b.data[i]
    }
    return t
}

func  GreaterEqual(a,b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v >= b.data[i]
    }
    return t
}

func Less(a, b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v < b.data[i]
    }
    return t
}

func  LessEqual(a, b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v <= b.data[i]
    }
    return t
}

func Equal(a, b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v == b.data[i]
    }
    return t
}

func NotEqual(a, b *Arrf) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v != b.data[i]
    }
    return t
}
