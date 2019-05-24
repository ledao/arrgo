package arrgo

import "fmt"

//ReShape 改变Tensor的形状，改变形状后的t将会返回。
//t本身的数据没有改动。
func (t *Tensor) ReShape(nShape ...int) *Tensor {
	if len(t.data) != GetShapeSize(nShape) {
		panic(SHAPE_ERROR)
	}

	tShape := make([]int, len(nShape))
	copy(tShape, nShape)
	t.shape = tShape
	t.strides = make([]int, len(t.shape)+1)
	for i := len(at.shape) - 1; i >= 0; i-- {
		t.strides[i] = t.strides[i+1] * t.shape[i]
	}
	return t
}

//两个多维数组形状相同，则返回true， 否则返回false。
func (a *Arrf) SameShapeTo(b *Arrf) bool {
	return SameIntSlice(a.shape, b.shape)
}

//将多个两维数组在垂直方向上组合起来，形成新的多维数组。
//不影响原多维数组。
func Vstack(arrs ...*Arrf) *Arrf {
	for i := range arrs {
		if arrs[i].Ndims() > 2 {
			fmt.Println("in Vstack function, array dimension cannot bigger than 2.")
			panic(SHAPE_ERROR)
		}
	}
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0].Copy()
	}

	return Concat(0, arrs...)
	//
	//var vlenSum int = 0
	//
	//var hlen int
	//if arrs[0].Ndims() == 1 {
	//	hlen = arrs[0].shape[0]
	//	vlenSum += 1
	//} else {
	//	hlen = arrs[0].shape[1]
	//	vlenSum += arrs[0].shape[0]
	//}
	//for i := 1; i < len(arrs); i++ {
	//	var nextHen int
	//	if arrs[i].Ndims() == 1 {
	//		nextHen = arrs[i].shape[0]
	//		vlenSum += 1
	//	} else {
	//		nextHen = arrs[i].shape[1]
	//		vlenSum += arrs[i].shape[0]
	//	}
	//	if hlen != nextHen {
	//		panic(SHAPE_ERROR)
	//	}
	//}
	//
	//data := make([]float64, vlenSum*hlen)
	//var offset = 0
	//for i := range arrs {
	//	copy(data[offset:], arrs[i].data)
	//	offset += len(arrs[i].data)
	//}
	//
	//return Array(data, vlenSum, hlen)
}

//将多个两维数组在水平方向上组合起来，形成新的多维数组。
//不影响原多维数组。
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

	return Concat(1, arrs...)

	//var hlenSum int = 0
	//var hBlockLens = make([]int, len(arrs))
	//var vlen int
	//if arrs[0].Ndims() == 1 {
	//	vlen = 1
	//	hlenSum += arrs[0].shape[0]
	//	hBlockLens[0] = arrs[0].shape[0]
	//} else {
	//	vlen = arrs[0].shape[0]
	//	hlenSum += arrs[0].shape[1]
	//	hBlockLens[0] = arrs[0].shape[1]
	//}
	//for i := 1; i < len(arrs); i++ {
	//	var nextVlen int
	//	if arrs[i].Ndims() == 1 {
	//		nextVlen = 1
	//		hlenSum += arrs[i].shape[0]
	//		hBlockLens[i] = arrs[i].shape[0]
	//	} else {
	//		nextVlen = arrs[i].shape[0]
	//		hlenSum += arrs[i].shape[1]
	//		hBlockLens[i] = arrs[i].shape[1]
	//	}
	//	if vlen != nextVlen {
	//		panic(SHAPE_ERROR)
	//	}
	//}
	//
	//data := make([]float64, hlenSum*vlen)
	//for i := 0; i < vlen; i++ {
	//	var curPos = 0
	//	for j := 0; j < len(arrs); j++ {
	//		copy(data[curPos+i*hlenSum:curPos+i*hlenSum+hBlockLens[j]], arrs[j].data[i*hBlockLens[j]:(i+1)*hBlockLens[j]])
	//		curPos += hBlockLens[j]
	//	}
	//}
	//
	//return Array(data, vlen, hlenSum)
}

//将一维数组扩充为二维
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

//将数组内部的元素铺平返回，创建新的数据副本。
func (a *Arrf) Flatten() *Arrf {
	ra := make([]float64, len(a.data))
	copy(ra, a.data)
	return Array(ra, len(a.data))
}
