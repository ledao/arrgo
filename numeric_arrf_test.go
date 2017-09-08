package arrgo

import "testing"
import "fmt"

func TestArrayCond1(t *testing.T) {
	arr := Array(nil)
	if SameFloat64Slice(arr.data, []float64{}) != true {
		t.Error("array data should be []float64{}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{0}) != true {
		t.Error("array shape should be []int{0}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{0, 1}) != true {
		t.Error("array strides should be []int{0, 1}, got ", arr.shape)
	}
}

func TestArrayCond2(t *testing.T) {
	arr := Array([]float64{1, 2, 3})
	if SameFloat64Slice(arr.data, []float64{1, 2, 3}) != true {
		t.Error("array data should be []float64{1,2,3}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{3}) != true {
		t.Error("array shape should be []int{3}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{3, 1}) != true {
		t.Error("array strides should be []int{3, 1}, got ", arr.shape)
	}
}

func TestArrayCond3ExceptionTwoNegtiveDims(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	Array([]float64{1, 2, 3, 4}, -1, -1, 4)
}

func TestArrayCond3ExceptionLengError(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	Array([]float64{1, 2, 3, 4}, 3, 4, 5)
}

func TestArrayCond3ExceptionDivError(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Exepcted shape error, got ", r)
		}
	}()

	Array([]float64{1, 2, 3, 4}, -1, 3)
}

func TestArrayCond3(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4}, 2, 2)
	if !SameIntSlice(arr.shape, []int{2, 2}) {
		t.Error("Expected [2, 2], got ", arr.shape)
	}
	if !SameIntSlice(arr.strides, []int{4, 2, 1}) {
		t.Error("Expected [4,2,1], got", arr.strides)
	}
	if !SameFloat64Slice(arr.data, []float64{1, 2, 3, 4}) {
		t.Error("Expected [1,2,3,4], got ", arr.data)
	}

	arr = Array([]float64{1, 2, 3, 4}, 2, -1)
	if !SameIntSlice(arr.shape, []int{2, 2}) {
		t.Error("Expected [2, 2], got ", arr.shape)
	}
	if !SameIntSlice(arr.strides, []int{4, 2, 1}) {
		t.Error("Expected [4,2,1], got", arr.strides)
	}
	if !SameFloat64Slice(arr.data, []float64{1, 2, 3, 4}) {
		t.Error("Expected [1,2,3,4], got ", arr.data)
	}
}

func TestArrayCond4(t *testing.T) {
	arr := Array(nil, 2, 3)
	if SameFloat64Slice(arr.data, []float64{0, 0, 0, 0, 0, 0}) != true {
		t.Error("array data should be []float64{0, 0, 0, 0, 0, 0}, got ", arr.data)
	}
	if SameIntSlice(arr.shape, []int{2, 3}) != true {
		t.Error("array shape should be []int{2, 3}, got ", arr.shape)
	}
	if SameIntSlice(arr.strides, []int{6, 3, 1}) != true {
		t.Error("array strides should be []int{6, 3, 1}, got ", arr.shape)
	}

	defer func() {
		err := recover()
		if err != SHAPE_ERROR {
			t.Error("should panic shape error, got ", err)
		}
	}()

	Array(nil, -1, 2, 3)
}

func TestArange(t *testing.T) {
	a1 := Arange(3)
	if !a1.Equal(Array([]float64{0, 1, 2})).All() {
		t.Error("Expected [0, 1, 2], got ", a1)
	}

	a1 = Arange(-3)
	if !a1.Equal(Array([]float64{0, -1, -2})).All() {
		t.Error("Expected [0, -1, -2], got ", a1)
	}

	a1 = Arange(1, 3)
	if !a1.Equal(Array([]float64{1, 2})).All() {
		t.Error("Expected [1,2], got ", a1)
	}

	a1 = Arange(-1, 2)
	if !a1.Equal(Array([]float64{-1, 0, 1})).All() {
		t.Error("Expected [-1, 0, 1], got ", a1)
	}

	a1 = Arange(2, -1)
	if !a1.Equal(Array([]float64{2, 1, 0})).All() {
		t.Error("Expected [2, 1, 0], got ", a1)
	}

	a1 = Arange(1, 4, 2)
	if !a1.Equal(Array([]float64{1, 3})).All() {
		t.Error("Expected [1, 3], got ", a1)
	}

	a1 = Arange(4, -1, -2)
	t.Log(a1)
	if !a1.Equal(Array([]float64{4, 2, 0})).All() {
		t.Error("Expected [4, 2, 0], got ", a1)
	}
}

func TestArangeIncrementExpection1(t *testing.T) {
	defer func() {
		r := recover()
		if r != PARAMETER_ERROR {
			t.Errorf("Expected PARAMTER ERROR, got ", r)
		}
	}()

	Arange(1, 3, -2)
}

func TestArangeIncrementExpection2(t *testing.T) {
	defer func() {
		r := recover()
		if r != PARAMETER_ERROR {
			t.Errorf("Expected PARAMTER ERROR, got ", r)
		}
	}()

	Arange(3, 1, 1)
}

func TestArangeNullParameterException(t *testing.T) {
	defer func() {
		r := recover()
		if r != PARAMETER_ERROR {
			t.Errorf("Expected PARAMETER ERROR, got ", r)
		}
	}()

	Arange()
}

func TestArrf_IsEmpty(t *testing.T) {
	empty := Array(nil)

	if empty.IsEmpty() != true {
		t.Errorf("Expected empty arra")
	}

	empty.data = make([]float64, 0)

	if empty.IsEmpty() != true {
		t.Errorf("Expected empty arra")
	}
}

func TestFull(t *testing.T) {
	arr := Full(1.0, 3)

	if !SameIntSlice(arr.shape, []int{3}) {
		t.Errorf("Expected [3], got ", arr.shape)
	}

	if !SameIntSlice(arr.strides, []int{3, 1}) {
		t.Errorf("Expected [3, 1], got ", arr.strides)
	}

	if !SameFloat64Slice(arr.data, []float64{1.0, 1.0, 1.0}) {
		t.Errorf("Expected [1.0, 1.0, 1.0], got ", arr.data)
	}
}

func TestFullException(t *testing.T) {
	defer func() {
		r := recover()

		if r != SHAPE_ERROR {
			t.Errorf("Expected SHAPE_ERROR, got ", r)
		}
	}()

	Full(1.0)
}

func TestArrf_Max(t *testing.T) {
	a := Arange(6).Reshape(2, 3)
	fmt.Println(a.Max())
	fmt.Println(a.Max(0))
	fmt.Println(a.Max(1))
	fmt.Println(a.Max(0, 1))
}

func TestArrf_Sort(t *testing.T) {
	a := Array([]float64{2, 3, 1, 5, 4, 1, 4, 5, 6, 4}).Reshape(2, 5)
	fmt.Println(a)
	a.Sort(1)
	fmt.Println(a)
}

//func TestVstack(t *testing.T) {
//	a := Arange(10)
//	b := Arange(10).Reshape(1, 10)
//	fmt.Println(Vstack(a, b))
//}
