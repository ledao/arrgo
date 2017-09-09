package arrgo

import (
    "fmt"
    "strings"
)

type Arrb struct {
    shape   []int
    strides []int
    data    []bool
}

//通过[]bool，形状来创建多维数组。
//输入参数1：data []bool，以·C· 顺序存储，作为多维数组的输入数据，内部复制一份新的internalData，不改变data。
//输入参数2：shape ...int，指定多维数组的形状，多维，类似numpy中的shape。
//	如果某一个（仅支持一个维度）维度为负数，则根据len(data)推断该维度的大小。
//情况1：如果不指定shape，而且data为nil，则创建一个空的*Arrb。
//情况2：如果不指定shape，而且data不为nil，则创建一个len(data)大小的一维*Arrb。
//情况3：如果指定shape，而且data不为nil，则根据data大小创建多维数组，如果len(data)不等于shape，或者len(data)不能整除shape，抛出异常。
//情况4：如果指定shape，而且data为nil，则创建shape大小的全为false的多维数组。
func ArrayB(data []bool, shape ...int) *Arrb {
	if len(shape) == 0 && data == nil {
		return &Arrb{
			shape:   []int{0},
			strides: []int{0, 1},
			data:    []bool{},
		}
	}

	if len(shape) == 0 && data != nil {
		internalData := make([]bool, len(data)) //复制data，不影响输入的值。
		copy(internalData, data)
		return &Arrb{
			shape:   []int{len(data)},
			strides: []int{len(data), 1},
			data:    internalData,
        }
    }

	if data == nil {
		for _, v := range shape {
			if v <= 0 {
				fmt.Println("shape should be positive when data is nill")
				panic(SHAPE_ERROR)
			}
		}
		length := ProductIntSlice(shape)
		internalShape := make([]int, len(shape))
		copy(internalShape, shape)
		strides := make([]int, len(shape)+1)
		strides[len(shape)] = 1
		for i := len(shape) - 1; i >= 0; i-- {
			strides[i] = strides[i+1] * internalShape[i]
		}

		return &Arrb{
			shape:   internalShape,
			strides: strides,
			data:    make([]bool, length),
		}
	}

	var dataLength = len(data)
	negativeIndex := -1
	internalShape := make([]int, len(shape))
	copy(internalShape, shape)
	for k, v := range shape {
		if v < 0 {
			if negativeIndex < 0 {
				negativeIndex = k
				internalShape[k] = 1
			} else {
				fmt.Println("shape can only have one negative demention.")
				panic(SHAPE_ERROR)
			}
		}
	}
	shapeLength := ProductIntSlice(internalShape)

	if dataLength < shapeLength {
		fmt.Println("data length is shorter than shape length.")
		panic(SHAPE_ERROR)
	}
	if (dataLength % shapeLength) != 0 {
		fmt.Println("data length cannot divided by shape length")
		panic(SHAPE_ERROR)
	}

	if negativeIndex >= 0 {
		internalShape[negativeIndex] = dataLength / shapeLength
	}

	strides := make([]int, len(internalShape)+1)
	strides[len(internalShape)] = 1
	for i := len(internalShape) - 1; i >= 0; i-- {
		strides[i] = strides[i+1] * internalShape[i]
	}

	internalData := make([]bool, len(data))
	copy(internalData, data)

	return &Arrb{
		shape:   internalShape,
		strides: strides,
		data:    internalData,
	}
}

//创建shape形状的多维布尔数组，全部填充为fillvalue。
//必须指定shape，否则抛出异常。
func FillB(fullValue bool, shape ...int) *Arrb {
	if len(shape) == 0 {
		fmt.Println("shape is empty!")
		panic(SHAPE_ERROR)
	}
	arr := ArrayB(nil, shape...)
	for i := range arr.data {
		arr.data[i] = fullValue
	}

	return arr
}

//创建全为false，形状位shape的多维布尔数组
func EmptyB(shape ...int) (a *Arrb) {
	a = FillB(false, shape...)
    return
}

func (a *Arrb) String() (s string) {
    switch {
    case a == nil:
        return "<nil>"
    case a.shape == nil || a.strides == nil || a.data == nil:
        return "<nil>"
    case a.strides[0] == 0:
        return "[]"
    }

    stride := a.strides[len(a.strides)-2]
    for i, k := 0, 0; i+stride <= len(a.data); i, k = i+stride, k+1 {

        t := ""
        for j, v := range a.strides {
            if i%v == 0 && j < len(a.strides)-2 {
                t += "["
            }
        }

        s += strings.Repeat(" ", len(a.shape)-len(t)-1) + t
        s += fmt.Sprint(a.data[i : i+stride])

        t = ""
        for j, v := range a.strides {
            if (i+stride)%v == 0 && j < len(a.strides)-2 {
                t += "]"
            }
        }

        s += t + strings.Repeat(" ", len(a.shape)-len(t)-1)
        if i+stride != len(a.data) {
            s += "\n"
            if len(t) > 0 {
                s += "\n"
            }
        }
    }
    return
}

//如果多维布尔数组元素都为真，返回true，否则返回false。
func (ab *Arrb) AllTrues() bool {
    if len(ab.data) == 0 {
        return false
    }
    for _, v := range ab.data {
        if v == false {
            return false
        }
    }
    return true
}

//如果多维布尔数组元素都为假，返回false，否则返回true。
func (ab *Arrb) AnyTrue() bool {
    if len(ab.data) == 0 {
        return false
    }
    for _, v := range ab.data {
        if v == true {
            return true
        }
    }
    return false
}

//返回多维数组中真值的个数。
func (a *Arrb) Sum() int {
    sum := 0
    for _, v := range a.data {
        if v {
            sum++
        }
    }
    return sum
}