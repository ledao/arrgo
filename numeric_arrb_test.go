package arrgo

import (
	"strings"
	"testing"
)

func TestArrayBCond1(t *testing.T) {
	arr := ArrayB(nil)
	if SameBoolSlice(arr.data, []bool{}) != true {
		t.Error("ArrayB data should be []bool{}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{0}) != true {
		t.Error("ArrayB shape should be []int{0}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{0, 1}) != true {
		t.Error("ArrayB strides should be []int{0, 1}, got ", arr.shape)
	}
}

func TestArrayBCond2(t *testing.T) {
	arr := ArrayB([]bool{true, true, true})
	if SameBoolSlice(arr.data, []bool{true, true, true}) != true {
		t.Error("ArrayB data should be []bool{true, true, true}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{3}) != true {
		t.Error("ArrayB shape should be []int{3}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{3, 1}) != true {
		t.Error("ArrayB strides should be []int{3, 1}, got ", arr.shape)
	}
}

func TestArrayBCond3ExceptionTwoNegtiveDims(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	ArrayB([]bool{true, true, true}, -1, -1, 4)
}

func TestArrayBCond3ExceptionLengError(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	ArrayB([]bool{true, true, true}, 3, 4, 5)
}

func TestArrayBCond3ExceptionDivError(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	ArrayB([]bool{true, true, true, true}, -1, 3)
}

func TestArrayBCond3(t *testing.T) {
	arr := ArrayB([]bool{true, true, true, true}, 2, 2)
	if !SameIntSlice(arr.shape, []int{2, 2}) {
		t.Error("Expected [true, true, true, true], got ", arr.shape)
	}
	if !SameIntSlice(arr.strides, []int{4, 2, 1}) {
		t.Error("Expected [4,2,1], got", arr.strides)
	}
	if !SameBoolSlice(arr.data, []bool{true, true, true, true}) {
		t.Error("Expected [true, true, true, true], got ", arr.data)
	}

	arr = ArrayB([]bool{true, true, true, true}, 2, -1)
	if !SameIntSlice(arr.shape, []int{2, 2}) {
		t.Error("Expected [2, 2], got ", arr.shape)
	}
	if !SameIntSlice(arr.strides, []int{4, 2, 1}) {
		t.Error("Expected [4,2,1], got", arr.strides)
	}
	if !SameBoolSlice(arr.data, []bool{true, true, true, true}) {
		t.Error("Expected [true, true, true, true], got ", arr.data)
	}
}

func TestArrayBCond4(t *testing.T) {
	arr := ArrayB(nil, 2, 3)
	if SameBoolSlice(arr.data, []bool{false, false, false, false, false, false}) != true {
		t.Error("ArrayB data should be []bool{false, false, false, false, false, false}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{2, 3}) != true {
		t.Error("ArrayB shape should be []int{2, 3}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{6, 3, 1}) != true {
		t.Error("ArrayB strides should be []int{6, 3, 1}, got ", arr.shape)
	}

	defer func() {
		err := recover()
		if err != SHAPE_ERROR {
			t.Error("should panic shape error, got ", err)
		}
	}()

	ArrayB(nil, -1, 2, 3)
}

func TestFillB(t *testing.T) {
	arr := FillB(true, 3)

	if !SameIntSlice(arr.shape, []int{3}) {
		t.Errorf("Expected [3], got %v", arr.shape)
	}

	if !SameIntSlice(arr.strides, []int{3, 1}) {
		t.Errorf("Expected [3, 1], got %v", arr.strides)
	}

	if !SameBoolSlice(arr.data, []bool{true, true, true}) {
		t.Errorf("Expected [true, true, true], got %v", arr.data)
	}
}

func TestFillBException(t *testing.T) {
	defer func() {
		r := recover()

		if r != SHAPE_ERROR {
			t.Errorf("Expected SHAPE_ERROR, got %v", r)
		}
	}()

	FillB(true)
}

func TestEmptyB(t *testing.T) {
	arr := EmptyB(3)
	if !SameBoolSlice(arr.data, []bool{false, false, false}) {
		t.Errorf("Expected [false, false, false], got %v", arr.data)
	}
}

func TestArrb_AllTrues(t *testing.T) {
	arr := ArrayB([]bool{true, true})
	if arr.AllTrues() != true {
		t.Errorf("Expected true, got %t", arr.AllTrues())
	}

	arr = ArrayB([]bool{true, false})
	if arr.AllTrues() != false {
		t.Errorf("EXepcted false, got %t", arr.AllTrues())
	}
}

func TestArrb_AnyTrue(t *testing.T) {
	arr := ArrayB([]bool{true, true})
	if arr.AnyTrue() != true {
		t.Errorf("Expected true, got %t", arr.AnyTrue())
	}

	arr = ArrayB([]bool{true, false})
	if arr.AnyTrue() != true {
		t.Errorf("EXepcted true, got %t", arr.AnyTrue())
	}

	arr = ArrayB([]bool{false, false})
	if arr.AnyTrue() != false {
		t.Errorf("EXepcted false, got %t", arr.AnyTrue())
	}
}

func TestArrb_String(t *testing.T) {
	var arr *Arrb

	if arr.String() != "<nil>" {
		t.Errorf("Expected <nil>, git %s", arr.String())
	}

	arr = EmptyB(2)
	arr.shape = nil
	if arr.String() != "<nil>" {
		t.Errorf("Expected <nil>, git %s", arr.String())
	}

	arr = EmptyB(2)
	arr.strides = make([]int, 2)
	if arr.String() != "[]" {
		t.Errorf("Expected [], got %s", arr.String())
	}

	arr = ArrayB([]bool{true, false}, 2, 1)
	if strings.Replace(arr.String(), "\n", ":", -1) != "[[true] : [false]]" {
		t.Errorf("Expected [[true]\n[false]], got %s", arr.String())
	}
}

func TestArrb_Sum(t *testing.T) {
	arr := ArrayB([]bool{true, true})
	if arr.Sum() != 2 {
		t.Errorf("Expected 2, got %d", arr.Sum())
	}

	arr = ArrayB([]bool{true, false})
	if arr.Sum() != 1 {
		t.Errorf("Expected 1, got %d", arr.Sum())
	}

	arr = ArrayB([]bool{false, false})
	if arr.Sum() != 0 {
		t.Errorf("Expected 0, got %d", arr.Sum())
	}
}
