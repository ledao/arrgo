package arrgo

import (
    "fmt"
    "strings"
)

type Arrgo struct {
    shape []int
    strides []int
    data []float64
}


func Array(data []float64, shape ...int)(a * Arrgo) {
    if len(shape) == 0 {
        switch {
        case data != nil:
            return &Arrgo{
                shape:   []int{len(data)},
                strides: []int{len(data), 1},
                data:    data,
            }
        default:
            return &Arrgo{
                shape:   []int{0},
                strides: []int{0, 0},
                data:    []float64{},
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

    a = &Arrgo{
        shape:   sh,
        strides: make([]int, len(shape)+1),
        data:    make([]float64, sz),
    }

    if data != nil {
        copy(a.data, data)
    }

    a.strides[len(shape)] = 1
    for i := len(shape) - 1; i >= 0; i-- {
        a.strides[i] = a.strides[i+1] * a.shape[i]
    }
    return
}

// Arange Creates an array in one of three different ways, depending on input:
//  Arange(stop):              Array64 from zero to stop
//  Arange(start, stop):       Array64 from start to stop(excluded), with increment of 1 or -1, depending on inputs
//  Arange(start, stop, step): Array64 from start to stop(excluded), with increment of step
//
// Any inputs beyond three values are ignored
func Arange(vals ...float64) (a *Arrgo) {
    var start, stop, step float64 = 0, 0, 1

    switch len(vals) {
    case 0:
        return empty(0)
    case 1:
        if vals[0] <= 0 {
            stop = -1
        } else {
            stop = vals[0] - 1
        }
    case 2:
        if vals[1] < vals[0] {
            step = -1
            stop = vals[1] + 1
        } else {
            stop = vals[1] - 1
        }
        start = vals[0]
    default:
        if vals[1] < vals[0] {
            stop = vals[1] + 1
        } else {
            stop = vals[1] -1
        }
        start, step = vals[0],  vals[2]
    }

    a = Array(nil, int((stop-start)/(step))+1)
    for i, v := 0, start; i < len(a.data); i, v = i+1, v+step {
        a.data[i] = v
    }
    return
}


// Internal function to create using the shape of another array
func empty(shape ...int) (a *Arrgo) {
    var sz int = 1
    for _, v := range shape {
        sz *= v
    }

    a = &Arrgo{
        shape:   shape,
        strides: make([]int, len(shape)+1),
        data:    make([]float64, sz),
    }

    a.strides[len(shape)] = 1
    for i := len(shape) - 1; i >= 0; i-- {
        a.strides[i] = a.strides[i+1] * a.shape[i]
    }
    return
}

//Return a new array of given shape and type, filled with ones.
//Parameters
//----------
//shape : int or sequence of ints
//Shape of the new array, e.g., ``(2, 3)`` or ``2``.
//dtype : data-type, optional
//The desired data-type for the array, e.g., `numpy.int8`.  Default is
//`numpy.float64`.
//order : {'C', 'F'}, optional
//Whether to store multidimensional data in C- or Fortran-contiguous
//(row- or column-wise) order in memory.
//Returns
//-------
//out : ndarray
//Array of ones with the given shape, dtype, and order.
func Ones(shape ...int) *Arrgo  {
    return Full(1, shape...)
}

//Return a new array of given shape and type, filled with `fill_value`.
//Parameters
//----------
//shape : int or sequence of ints
//Shape of the new array, e.g., ``(2, 3)`` or ``2``.
//fill_value : scalar
//Fill value.
//dtype : data-type, optional
//The desired data-type for the array, e.g., `np.int8`.  Default
//is `float`, but will change to `np.array(fill_value).dtype` in a
//future release.
//order : {'C', 'F'}, optional
//Whether to store multidimensional data in C- or Fortran-contiguous
//(row- or column-wise) order in memory.
//Returns
//out : ndarray
//Array of `fill_value` with the given shape, dtype, and order.
func Full(fullValue float64, shape ...int) *Arrgo {
    arr := Array(nil, shape...)
    if fullValue == 0 {
        return arr
    }
    return arr.AddC(fullValue)
}

// String Satisfies the Stringer interface for fmt package
func (a *Arrgo) String() (s string) {
    switch {
    case a == nil:
        return "<nil>"
    case a.data == nil || a.shape == nil || a.strides == nil:
        return "<nil>"
    case a.strides[0] == 0:
        return "[]"
    case len(a.shape) == 1:
        return fmt.Sprint(a.data)
    }

    stride := a.shape[len(a.shape)-1]

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

func(a *Arrgo) At(index ...int) float64 {
    idx , err := a.valIndex(index...)
    if err != nil {
        panic(err)
    }
    return a.data[idx]
}

func(a *Arrgo) valIndex(index ...int) (int, error) {
    idx := 0
    if len(index) > len(a.shape) {
        return -1, INDEX_ERROR
    }
    for i, v := range index {
        if v >= a.shape[i] || v < 0 {
            return -1, INDEX_ERROR
        }
        idx += v * a.strides[i+1]
    }
    return idx, nil
}

// Reshape Changes the size of the array axes.  Values are not changed or moved.
// This must not change the size of the array.
// Incorrect dimensions will return a nil pointer
func (a *Arrgo) Reshape(shape ...int) *Arrgo {
    if len(shape) == 0 {
        return a
    }

    var sz = 1
    sh := make([]int, len(shape))
    for _, v := range shape {
        if v < 0 {
            panic(SHAPE_ERROR)
        }
        sz *= v
    }
    copy(sh, shape)

    if sz != len(a.data) {
        panic(SHAPE_ERROR)
    }

    a.strides = make([]int, len(sh)+1)
    tmp := 1
    for i := len(a.strides) - 1; i > 0; i-- {
        a.strides[i] = tmp
        tmp *= sh[i-1]
    }
    a.strides[0] = tmp
    a.shape = sh

    return a
}

