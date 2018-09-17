package arrgo

import "testing"

func TestArrf_Reshape(t *testing.T) {
	arr := Array([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	arr2 := arr.Reshape(3, 2)

	if !SameIntSlice(arr.strides, []int{6, 2, 1}) {
		t.Errorf("Expected [6,2,1], got ", arr2.strides)
	}
	if !SameIntSlice(arr.shape, []int{3, 2}) {
		t.Errorf("Expected [3, 2], got ", arr.shape)
	}
	if !SameIntSlice(arr2.shape, []int{3, 2}) {
		t.Errorf("Expected [3, 2], got ", arr2.shape)
	}
}

func TestArrf_ReshapeException(t *testing.T) {
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got ", r)
		}
	}()

	Arange(4).Reshape(5)
}

func TestArrf_SameShapeTo(t *testing.T) {
	a := Arange(4).Reshape(2, 2)
	b := Array([]float64{3, 4, 5, 6}, 2, 2)
	if a.SameShapeTo(b) != true {
		t.Errorf("Expected true, got %d", a.SameShapeTo(b))
	}
}

func TestVstack(t *testing.T) {
	if Vstack() != nil {
		t.Errorf("Expected nil, got %s", Vstack())
	}

	a := Arange(3)
	stacked := Vstack(a)
	if !stacked.Equal(Arange(3)).AllTrues() {
		t.Errorf("Expected [0, 1, 2], got %s", stacked)
	}

	b := Array([]float64{3, 4, 5})
	stacked = Vstack(a, b)
	if !stacked.Equal(Array([]float64{0, 1, 2, 3, 4, 5}, 2, 3)).AllTrues() {
		t.Errorf("Expected [[0 1 2] [3 4 5]], got %s", stacked)
	}

	a = Arange(2)
	b = Arange(4).Reshape(2, 2)
	stacked = Vstack(a, b)
	if !stacked.Equal(Array([]float64{0, 1, 0, 1, 2, 3}, 3, 2)).AllTrues() {
		t.Errorf("Expected [[0,1], [0,1], [2, 3]], got %s", stacked)
	}
}

func TestVstackException(t *testing.T) {
	a := Arange(4).Reshape(1, 2, 2)
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got %s", r)
		}
	}()

	Vstack(a)
}

func TestVstackException2(t *testing.T) {
	a := Arange(4)
	b := Arange(5)
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got ", r)
		}
	}()

	Vstack(a, b)
}

func TestHstack(t *testing.T) {
	if Hstack() != nil {
		t.Errorf("Expected nil, got ", Hstack())
	}

	a := Arange(3)
	stacked := Hstack(a)
	if !stacked.Equal(Arange(3)).AllTrues() {
		t.Errorf("Expected [0, 1, 2], got ", stacked)
	}
	a = a.Reshape(3, 1)
	b := Array([]float64{3, 4, 5}).Reshape(3, 1)
	stacked = Hstack(a, b)
	if !stacked.Equal(Array([]float64{0, 3, 1, 4, 2, 5}, 3, 2)).AllTrues() {
		t.Errorf("Expected [[0 3] [1 4], [2 5]], got ", stacked)
	}

	a = Arange(2).Reshape(2, 1)
	b = Arange(4).Reshape(2, 2)
	stacked = Hstack(a, b)
	if !stacked.Equal(Array([]float64{0, 0, 1, 1, 2, 3}, 2, 3)).AllTrues() {
		t.Errorf("Expected [[0, 0, 1], [1, 2, 3]], got ", stacked)
	}
}

func TestHstackException(t *testing.T) {
	a := Arange(4).Reshape(1, 2, 2)
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got ", r)
		}
	}()

	Hstack(a)
}

func TestHstackException2(t *testing.T) {
	a := Arange(4).Reshape(4, 1)
	b := Arange(5).Reshape(5, 1)
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got ", r)
		}
	}()

	Hstack(a, b)
}

func TestConcat(t *testing.T) {
	if Concat(0) != nil {
		t.Errorf("Expected nil, got ", Concat(0))
	}
	concated := Concat(0, Arange(2))
	if !concated.Equal(Arange(2)).AllTrues() {
		t.Errorf("Expected [0, 1], got ", concated)
	}

	a := Arange(3)
	b := Arange(1, 4)

	concated = Concat(0, a, b)
	if !concated.Equal(Array([]float64{0, 1, 2, 1, 2, 3}, 2, 3)).AllTrues() {
		t.Errorf("Expected [[0,1,2], [1,2,3]], got ", concated)
	}

	a = Arange(3)
	b = Arange(1, 4)

	concated = Concat(1, a, b)
	t.Log(concated)
	if !concated.Equal(Array([]float64{0, 1, 2, 1, 2, 3}, 1, 6)).AllTrues() {
		t.Errorf("Expected [[0,1,2,1,2,3]], got ", concated)
	}

}

func TestConcatException(t *testing.T) {
	a := Arange(4)
	b := Arange(1, 4)

	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Errorf("Expected shape error, got ", r)
		}
	}()

	Concat(0, a, b)
}

func TestConcatException2(t *testing.T) {
	a := Arange(4)
	b := Arange(1, 4)

	defer func() {
		r := recover()
		if r != PARAMETER_ERROR {
			t.Errorf("Expected PARAMETER_ERROR, got ", r)
		}
	}()

	Concat(2, a, b)
}

func TestAtLeast2D(t *testing.T) {
	a := Arange(10)
	AtLeast2D(a)
	if !SameIntSlice(a.shape, []int{1, 10}) {
		t.Error("Expected [1, 10], got ", a.shape)
	}

	a.Reshape(1, 1, 10)
	AtLeast2D(a)
	if !SameIntSlice(a.shape, []int{1, 1, 10}) {
		t.Error("Expected [1, 1, 10], got ", a.shape)
	}

	if AtLeast2D(nil) != nil {
		t.Error("Expected nil, got ", AtLeast2D(nil))
	}
}

func TestAtLeast2D2(t *testing.T) {
	if AtLeast2D(nil) != nil {
		t.Errorf("Expected nil, got ", AtLeast2D(nil))
	}

	arr := Arange(3)
	AtLeast2D(arr)

	if !SameIntSlice(arr.shape, []int{1, 3}) {
		t.Errorf("expected true, got false")
	}

	arr = Arange(3).Reshape(3, 1)
	AtLeast2D(arr)

	if !SameIntSlice(arr.shape, []int{3, 1}) {
		t.Errorf("expected true, got false")
	}
}

func TestArrf_Flatten(t *testing.T) {
	arr := Arange(3).Reshape(3, 1)
	flattened := arr.Flatten()

	if !flattened.SameShapeTo(Arange(3)) {
		t.Errorf("expected [3], got ", flattened.shape)
	}
}
