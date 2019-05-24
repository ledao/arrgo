package arrgo

import "fmt"

//Cat 将多个多维数组在指定的轴上组合起来。
func Cat(dim int, arrs ...*Arrf) *Arrf {
	if len(arrs) == 0 {
		panic(EMPTY_ARRAY_ERROR)
	}

	if len(arrs) == 1 {
		return Copy(arrs[0])
	}

	if dim >= arrs[0].Ndims() {
		fmt.Println("dim is bigger than dimensions num.")
		panic(PARAMETER_ERROR)
	}

	var newShape = make([]int, arrs[0].Ndims())
	for index, firstL := range arrs[0].shape {
		if index == dim {
			newShape[index] += firstL
			for j := 1; j < len(arrs); j++ {
				newShape[index] += arrs[j].shape[index]
			}
		} else {
			newShape[index] = firstL
			for j := 1; j < len(arrs); j++ {
				if firstL != arrs[j].shape[index] {
					fmt.Println("Tensor ", j, " ", index, " dim not same to others, Cat fails.")
					panic(SHAPE_ERROR)
				}
			}
		}
	}

	var times = 0
	if dim == 0 {
		times = 1
	} else {
		times = GetShapeSize(arrs[0].shape[0:dim])
	}

	var data = make([]float64, GetShapeSize(newShape))

	var curPos = 0
	for i := 0; i < times; i++ {
		for j := 0; j < len(arrs); j++ {
			var l = GetShapeSize(arrs[j].shape[dim:])
			copy(data[curPos:curPos+l], arrs[j].data[i*l:(i+1)*l])
			curPos += l
		}
	}

	return NewTensor(data, newShape...)
}

func Chunk(tensor *Tensor, chunks int, dim int) []*Tensor {

}
