package arrgo

import (
	"testing"
)

func TestSameIntSlice(t *testing.T) {
	var s1 []int = nil
	var s2 []int = nil
	if SameIntSlice(s1, s2) != false {
		t.Error("if one of args is nil, the result should be false, but go ", SameIntSlice(s1, s2))
	}
	s3 := []int{1, 2, 3}
	s4 := []int{1, 2}
	if SameIntSlice(s3, s4) != false {
		t.Error("different length should get false, got ", SameIntSlice(s3, s4))
	}
	s5 := []int{1, 2, 4}
	if SameIntSlice(s3, s5) != false {
		t.Error("bit wise different should get false, got ", SameIntSlice(s3, s5))
	}
	s6 := []int{1, 2, 3}
	if SameIntSlice(s3, s6) != true {
		t.Error("same int[] should get true, got ", SameIntSlice(s3, s6))
	}
}
