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
