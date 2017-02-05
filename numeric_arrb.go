package arrgo

import (
    "fmt"
    "strings"
)

type Arrb struct {
    shape   []int
    strides []int
    data    []bool
}


func ArrayB(data []bool, shape ...int) (a *Arrb) {
    if len(shape) == 0 {
        switch {
        case data != nil:
            dataCopy := make([]bool, len(data))
            copy(dataCopy, data)
            return &Arrb{
                shape:   []int{len(data)},
                strides: []int{len(data), 1},
                data:    dataCopy,
            }
        default:
            return &Arrb{
                shape:   []int{0},
                strides: []int{0, 0},
                data:    []bool{},
            }
        }
    }

    var sz = 1
    sh := make([]int, len(shape))
    for _, v := range shape {
        if v < 0 {
            return
        }
        sz *= v
    }
    copy(sh, shape)

    a = &Arrb{
        shape:   sh,
        strides: make([]int, len(shape) + 1),
        data:    make([]bool, sz),
    }

    if data != nil {
        copy(a.data, data)
    }

    a.strides[len(shape)] = 1
    for i := len(shape) - 1; i >= 0; i-- {
        a.strides[i] = a.strides[i + 1] * a.shape[i]
    }
    return
}

func EmptyB(shape ...int) (a *Arrb) {
    var sz int = 1
    for _, v := range shape {
        sz *= v
    }
    shapeCopy := make([]int, len(shape))
    copy(shapeCopy, shape)
    a = &Arrb{
        shape:   shapeCopy,
        strides: make([]int, len(shape) + 1),
        data:    make([]bool, sz),
    }

    a.strides[len(shape)] = 1
    for i := len(shape) - 1; i >= 0; i-- {
        a.strides[i] = a.strides[i + 1] * a.shape[i]
    }
    return
}


func FullB(value bool, shape ...int) *Arrb {
    a := EmptyB(shape...)
    for i := range a.data {
        a.data[i] = value
    }
    return a
}

func (a *Arrb) String() (s string) {
    switch {
    case a == nil:
        return "<nil>"
    case a.shape == nil || a.strides == nil || a.data == nil:
        return "<nil>"
    case a.strides[0] == 0:
        return "[]"
    }

    stride := a.strides[len(a.strides)-2]
    for i, k := 0, 0; i+stride <= len(a.data); i, k = i+stride, k+1 {

        t := ""
        for j, v := range a.strides {
            if i%v == 0 && j < len(a.strides)-2 {
                t += "["
            }
        }

        s += strings.Repeat(" ", len(a.shape)-len(t)-1) + t
        s += fmt.Sprint(a.data[i : i+stride])

        t = ""
        for j, v := range a.strides {
            if (i+stride)%v == 0 && j < len(a.strides)-2 {
                t += "]"
            }
        }

        s += t + strings.Repeat(" ", len(a.shape)-len(t)-1)
        if i+stride != len(a.data) {
            s += "\n"
            if len(t) > 0 {
                s += "\n"
            }
        }
    }
    return
}

func(ab *Arrb) All() bool {
    for _, v := range ab.data {
        if v == false {
            return false
        }
    }
    return true
}

func(ab *Arrb) Any() bool {
    for _, v := range ab.data {
        if v == true {
            return true
        }
    }
    return false
}
