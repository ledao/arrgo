package arrgo

import "fmt"

//改变原始多维数组的形状，并返回改变后的多维数组的指引引用。
//不会创建新的数据副本。
//如果新的shape的大小和原来多维数组的大小不同，则抛出异常。
func (a *Arrf) Reshape(shape ...int) *Arrf {
	if a.Length() != ProductIntSlice(shape) {
		fmt.Println("new shape length does not equal to original array length.")
		panic(SHAPE_ERROR)
	}

	internalShape := make([]int, len(shape))
	copy(internalShape, shape)
	a.shape = internalShape

	a.strides = make([]int, len(a.shape)+1)
	a.strides[len(a.shape)] = 1
	for i := len(a.shape) - 1; i >= 0; i-- {
		a.strides[i] = a.strides[i+1] * a.shape[i]
	}

	return a
}

func (a *Arrf) SameShapeTo(b *Arrf) bool {
	return SameIntSlice(a.shape, b.shape)
}

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

	data := make([]float64, vlenSum*hlen)
	var offset = 0
	for i := range arrs {
		copy(data[offset:], arrs[i].data)
		offset += len(arrs[i].data)
	}

	return Array(data, vlenSum, hlen)
}

func Hstack(arrs ...*Arrf) *Arrf {
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

	var hlenSum int = 0
	var hBlockLens = make([]int, len(arrs))
	var vlen int
	if arrs[0].Ndims() == 1 {
		vlen = 1
		hlenSum += arrs[0].shape[0]
		hBlockLens[0] = arrs[0].shape[0]
	} else {
		vlen = arrs[0].shape[0]
		hlenSum += arrs[0].shape[1]
		hBlockLens[0] = arrs[0].shape[1]
	}
	for i := 1; i < len(arrs); i++ {
		var nextVlen int
		if arrs[i].Ndims() == 1 {
			nextVlen = 1
			hlenSum += arrs[i].shape[0]
			hBlockLens[i] = arrs[i].shape[0]
		} else {
			nextVlen = arrs[i].shape[0]
			hlenSum += arrs[i].shape[1]
			hBlockLens[i] = arrs[i].shape[1]
		}
		if vlen != nextVlen {
			panic(SHAPE_ERROR)
		}
	}

	data := make([]float64, hlenSum*vlen)
	for i := 0; i < vlen; i++ {
		var curPos = 0
		for j := 0; j < len(arrs); j++ {
			copy(data[curPos+i*hlenSum:curPos+i*hlenSum+hBlockLens[j]], arrs[j].data[i*hBlockLens[j]:(i+1)*hBlockLens[j]])
			curPos += hBlockLens[j]
		}
	}

	return Array(data, vlen, hlenSum)
}

func Concat(axis int, arrs ...*Arrf) *Arrf {
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0].Copy()
	}

	var newShape = make([]int, arrs[0].Ndims())
	for index, firstL := range arrs[0].shape {
		if index == axis {
			newShape[index] += firstL
			for j := 1; j < len(arrs); j++ {
				newShape[index] += arrs[j].shape[index]
			}
		} else {
			newShape[index] = firstL
			for j := 1; j < len(arrs); j++ {
				if firstL != arrs[j].shape[index] {
					panic(SHAPE_ERROR)
				}
			}
		}
	}

	var times = 0
	if axis == 0 {
		times = 1
	} else {
		times = ProductIntSlice(arrs[0].shape[0:axis])
	}

	var data = make([]float64, ProductIntSlice(newShape))

	var curPos = 0
	for i := 0; i < times; i++ {
		for j := 0; j < len(arrs); j++ {
			var l = ProductIntSlice(arrs[j].shape[axis:])
			copy(data[curPos:curPos+l], arrs[j].data[i*l:(i+1)*l])
			curPos += l
		}
	}

	return Array(data, newShape...)
}

func AtLeast2D(a *Arrf) *Arrf {
	if a == nil {
		return nil
	} else if a.Ndims() >= 2 {
		return a
	} else {
		newShpae := make([]int, 2)
		newShpae[0] = 1
		newShpae[1] = a.shape[0]
		a.shape = newShpae
		return a
	}
}

//
func (a *Arrf) Flatten() *Arrf {
	ra := make([]float64, len(a.data))
	copy(ra, a.data)
	return Array(ra, len(a.data))
}
