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

func TestHstack(t *testing.T) {
	var a = Arange(10)
	var b = Arange(10)
	t.Log(Hstack(a, b, a, b))
}

func TestConcat(t *testing.T) {
	var a = Arange(1, 11).Reshape(2, 5)
	var b = Arange(10).Reshape(2, 5)
	t.Log(Concat(0, a, b))
	t.Log(Concat(1, a, b))
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
