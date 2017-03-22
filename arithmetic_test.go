package arrgo

import (
    "testing"
)

func TestArrf_AddC(t *testing.T) {
    arr := Arange(0, 10, 2)
    add := arr.AddC(2)
    if !add.Equal(Array([]float64{2,4,6,8,10})).All() {
        t.Error("Expected [2,4,6,8,10], got ", add)
    }
}

func TestArrf_Add(t *testing.T) {
    var a = Array([]float64{1,2,3,4,5,6}, 2, 3)
    var b = Array([]float64{6,5,4,3,2,1}, 2, 3)
    var c = a.Add(b)
    if !c.Equal(Full(7, 2, 3)).All() {
        t.Error("Expected [[7,7,7],[7,7,7]], got ", c)
    }
}

func TestArrf_Add_NilException(t *testing.T) {
    var a = Array([]float64{1,2,3,4,5,6}, 2, 3)

    defer func(){
       var rec = recover()
        if rec != SHAPE_ERROR {
            t.Error("Expected SHAPE ERROR, got ", rec)
        }
    }()
    a.Add(nil)
}

func TestArrf_Add_NDimException(t *testing.T) {
    var a = Array([]float64{1,2,3,4,5,6})
    var b = Array([]float64{1,2,3}, 3, 1)
    defer func(){
       var rec = recover()
        if rec != SHAPE_ERROR {
            t.Error("Expected SHAPE ERROR, got ", rec)
        }
    }()
    a.Add(b)
}
