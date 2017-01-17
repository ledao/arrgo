package arrgo

import "testing"

func TestArray(t *testing.T) {
    arr := Arange(-10)
    t.Error("log: ", arr)
}
