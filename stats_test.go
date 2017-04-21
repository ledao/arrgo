package arrgo

import "testing"

func TestSum2(t *testing.T) {
	var arr = Arange(100).Reshape(2, 5, 10)
	t.Log(arr.Sum(0, 1, 2))
}
