package arrgo

import (
	"testing"
)

func TestEqualEmptyArrayException(t *testing.T) {
	a := Array(nil)
	b := Array(nil)
	defer func() {
		r := recover()
		if r != EMPTY_ARRAY_ERROR {
			t.Error("Expected EMPTY_ARRAY_ERROR, got ", r)
		}
	}()
	a.Equal(b)
}

func TestEqualShapeException(t *testing.T) {
	a := Array(nil, 3,4)
	b := Array(nil, 1,2)
	defer func() {
		r := recover()
		if r != SHAPE_ERROR {
			t.Error("Expected SHAPE_ERROR, got ", r)
		}
	}()
	a.Equal(b)
}

func TestEqual(t *testing.T) {
	a := Array([]float64{1,2,3}, )
	b := Array([]float64{1,2,4}, )

	var compares = a.Equal(b)
	if compares.data[2] != false {
		t.Error("Expected [true, true, false], got ", compares)
	}
}
