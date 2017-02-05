
package arrgo


func (a *Arrb) LogicalAnd(b *Arrb) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v && b.data[i]
    }
    return t
}

func (a *Arrb) LogicalOr(b *Arrb) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v || b.data[i]
    }
    return t
}

func (a *Arrb) LogicalNot() *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = !v
    }
    return t
}

func  LogicalAnd(a, b *Arrb) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v && b.data[i]
    }
    return t
}

func LogicalOr(a, b *Arrb) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = v || b.data[i]
    }
    return t
}

func LogicalNot(a *Arrb) *Arrb {
    var t = EmptyB(a.shape...)
    for i, v := range a.data {
        t.data[i] = !v
    }
    return t
}