package arrgo

import "testing"
import "fmt"

func TestArray(t *testing.T) {
	arr := Arange(-10)
	t.Log("log: ", arr)
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

func TestVstack(t *testing.T) {
	a := Arange(10)
	b := Arange(10).Reshape(1, 10)
	fmt.Println(Vstack(a, b))
}
