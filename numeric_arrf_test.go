package arrgo

import "testing"

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
	if !a1.Equal(Array([]float64{0, 1, 2})).AllTrues() {
		t.Error("Expected [0, 1, 2], got ", a1)
	}

	a1 = Arange(-3)
	if !a1.Equal(Array([]float64{0, -1, -2})).AllTrues() {
		t.Error("Expected [0, -1, -2], got ", a1)
	}

	a1 = Arange(1, 3)
	if !a1.Equal(Array([]float64{1, 2})).AllTrues() {
		t.Error("Expected [1,2], got ", a1)
	}

	a1 = Arange(-1, 2)
	if !a1.Equal(Array([]float64{-1, 0, 1})).AllTrues() {
		t.Error("Expected [-1, 0, 1], got ", a1)
	}

	a1 = Arange(2, -1)
	if !a1.Equal(Array([]float64{2, 1, 0})).AllTrues() {
		t.Error("Expected [2, 1, 0], got ", a1)
	}

	a1 = Arange(1, 4, 2)
	if !a1.Equal(Array([]float64{1, 3})).AllTrues() {
		t.Error("Expected [1, 3], got ", a1)
	}

	a1 = Arange(4, -1, -2)
	if !a1.Equal(Array([]float64{4, 2, 0})).AllTrues() {
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

func TestOnes(t *testing.T) {
	arr := Ones(3)

	if !SameIntSlice(arr.shape, []int{3}) {
		t.Error("Expected [3], got ", arr.shape)
	}

	if !SameIntSlice(arr.strides, []int{3, 1}) {
		t.Errorf("Expected [3, 1], got ", arr.strides)
	}

	if !SameFloat64Slice(arr.data, []float64{1, 1, 1}) {
		t.Errorf("Expected [1, 1, 1], got ", arr.data)
	}
}

func TestZeros(t *testing.T) {
	arr := Zeros(3)

	if !SameIntSlice(arr.shape, []int{3}) {
		t.Error("Expected [3], got ", arr.shape)
	}

	if !SameIntSlice(arr.strides, []int{3, 1}) {
		t.Errorf("Expected [3, 1], got ", arr.strides)
	}

	if !SameFloat64Slice(arr.data, []float64{0, 0, 0}) {
		t.Errorf("Expected [0,0,0], got ", arr.data)
	}
}

func TestArrf_At(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)

	if arr.At(0, 1) != 2.0 {
		t.Errorf("Expected 2.0, got ", arr.At(1, 0))
	}

	if arr.At(0) != 1.0 {
		t.Errorf("Expected 1.0, got ", arr.At(0))
	}

	if arr.At(1) != 4.0 {
		t.Errorf("Expected 4.0, got ", arr.At(1.0))
	}
}

func TestArrf_AtLongIndexException(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)

	defer func() {
		r := recover()
		if r != INDEX_ERROR {
			t.Errorf("Expected INDEX_ERROR, got ", r)
		}
	}()

	arr.At(0, 0, 1)
}

func TestArrf_AtIndexOutofRangeException(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)

	defer func() {
		r := recover()
		if r != INDEX_ERROR {
			t.Errorf("Expected INDEX_ERROR, got ", r)
		}
	}()

	arr.At(2, 0)
}

func TestArrf_ValIndex(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)

	index := arr.valIndex(0, 1)
	if index != 1 {
		t.Errorf("Expected 1, got ", index)
	}

	index = arr.valIndex(0)
	if index != 0 {
		t.Errorf("Expected 0, got ", index)
	}

	index = arr.valIndex(1)
	if index != 3 {
		t.Errorf("Expected 3, got ", index)
	}
}

func TestArrf_ValIndexExpection1(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	defer func() {
		r := recover()
		if r != INDEX_ERROR {
			t.Errorf("Expected INDEX_ERROR, got ", r)
		}
	}()
	arr.valIndex(0, 1, 0)
}

func TestArrf_ValIndexExpection2(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	defer func() {
		r := recover()
		if r != INDEX_ERROR {
			t.Errorf("Expected INDEX_ERROR, got ", r)
		}
	}()
	arr.valIndex(2)
}

func TestArrf_Length(t *testing.T) {
	arr := Array(nil, 2, 3)

	if arr.Length() != 6 {
		t.Errorf("Expected 6, got ", arr.Length())
	}
}

func TestEye(t *testing.T) {
	arr := Eye(2)

	if !arr.Equal(Array([]float64{1, 0, 0, 1}, 2, 2)).AllTrues() {
		t.Errorf("Expected [1, 0, 0, 1], got ", arr)
	}

	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected SHAPE_ERROR, got ", r)
		}
	}()

	Eye(0)
}

func TestArrf_Set(t *testing.T) {
	arr := Zeros(3)
	arr.Set(10, 1)

	if arr.Get(1) != 10 {
		t.Errorf("Expected 10, got ", arr.Get(10))
	}
}

func TestArrf_Values(t *testing.T) {
	arr := Array([]float64{1, 2, 3})

	values := arr.Values()

	if !SameFloat64Slice(values, []float64{1, 2, 3}) {
		t.Errorf("Expected [1.0, 2.0, 3.0], got ", values)
	}
	values[0] = 100

	if arr.data[0] != 100 {
		t.Errorf("Expected 100, got ", arr.data[0])
	}
}

func TestLinspace(t *testing.T) {
	arr := Linspace(1, 2, 5)

	if !arr.Equal(Array([]float64{1, 1.25, 1.5, 1.75, 2})).AllTrues() {
		t.Errorf("Expected [1, 1.25, 1.5, 1.75, 2], got ", arr)
	}

	arr = Linspace(2, 1, 5)

	if !arr.Equal(Array([]float64{2, 1.75, 1.5, 1.25, 1})).AllTrues() {
		t.Errorf("Expected [2, 1.75, 1.5, 1.25, 1], got ", arr)
	}

	arr = Linspace(-2, -1, 5)

	if !arr.Equal(Array([]float64{-2, -1.75, -1.5, -1.25, -1})).AllTrues() {
		t.Errorf("Expected [-2, -1.75, -1.5, -1.25, -1], got ", arr)
	}

	arr = Linspace(-1, -2, 5)

	if !arr.Equal(Array([]float64{-1, -1.25, -1.5, -1.75, -2})).AllTrues() {
		t.Errorf("Expected [-1, -1.25, -1.5, -1.75, -2], got ", arr)
	}

	arr = Linspace(-1, 2, 5)

	if !arr.Equal(Array([]float64{-1, -0.25, 0.5, 1.25, 2})).AllTrues() {
		t.Errorf("Expected [-1, -0.25, 0.5, 1.25, 2], got ", arr)
	}
}

func TestArrf_Copy(t *testing.T) {
	arr := Ones(2)
	arrCopy := arr.Copy()
	arr.Set(10, 0)

	if !arrCopy.Equal(Array([]float64{1, 1})).AllTrues() {
		t.Errorf("Expected [1, 1], got ", arrCopy)
	}
}

func TestArrf_Ndims(t *testing.T) {
	arr := Arange(10)
	if arr.Ndims() != 1 {
		t.Errorf("Expected 1, got ", arr.Ndims())
	}

	arr.Reshape(2, 5)
	if arr.Ndims() != 2 {
		t.Errorf("Expected 2, got ", arr.Ndims())
	}

	arr.Reshape(2, 5, 1)
	if arr.Ndims() != 3 {
		t.Errorf("Expected 3, got ", arr.Ndims())
	}
}

func TestArrf_Transpose(t *testing.T) {
	arr := Arange(4).Reshape(2, 2)

	if !arr.Equal(Array([]float64{0, 1, 2, 3}, 2, 2)).AllTrues() {
		t.Errorf("Expected [[0,1],[2,3]], got ", arr)
	}

	arrTransposed := arr.Transpose()
	if !arrTransposed.Equal(Array([]float64{0, 2, 1, 3}, 2, 2)).AllTrues() {
		t.Errorf("Expected [[0,2,], [1,3]], got ", arrTransposed)
	}

	arrTransposed = arr.Transpose(1, 0)
	if !arrTransposed.Equal(Array([]float64{0, 2, 1, 3}, 2, 2)).AllTrues() {
		t.Errorf("Expected [[0,2,], [1,3]], got ", arrTransposed)
	}
}

func TestArrf_TransposeException(t *testing.T) {
	arr := Arange(4)

	defer func() {
		r := recover()
		if r != DIMENTION_ERROR {
			t.Errorf("Expected DIMENTION_ERROR, got ", r)
		}
	}()
	arr.Transpose(0, 1)
}