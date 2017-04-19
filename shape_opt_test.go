package arrgo

import "testing"

func TestHstack(t *testing.T) {
	var a = Arange(10)
	var b = Arange(10)
	t.Log(Hstack(a, b, a, b))
}
