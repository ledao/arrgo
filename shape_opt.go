package arrgo

func Vstack(arrs ...*Arrf) *Arrf {
	for i := range arrs {
		if arrs[i].Ndims() > 2 {
			panic(SHAPE_ERROR)
		}
	}
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0].Copy()
	}

	var vlenSum int = 0

	var hlen int
	if arrs[0].Ndims() == 1 {
		hlen = arrs[0].shape[0]
		vlenSum += 1
	} else {
		hlen = arrs[0].shape[1]
		vlenSum += arrs[0].shape[0]
	}
	for i := 1; i < len(arrs); i++ {
		var nextHen int
		if arrs[i].Ndims() == 1 {
			nextHen = arrs[i].shape[0]
			vlenSum += 1
		} else {
			nextHen = arrs[i].shape[1]
			vlenSum += arrs[i].shape[0]
		}
		if hlen != nextHen {
			panic(SHAPE_ERROR)
		}
	}

	data := make([]float64, vlenSum * hlen)
	var offset = 0
	for i := range arrs {
		copy(data[offset:], arrs[i].data)
		offset += len(arrs[i].data)
	}

	return Array(data, vlenSum, hlen)
}
