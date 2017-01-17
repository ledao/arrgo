//+build !amd64 noasm appengine

package asm

func findBool(vals []bool, find bool) (flg bool) {
	for _, v := range vals {
		if v == find {
			return true
		}
	}
	return false
}
