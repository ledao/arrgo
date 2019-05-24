package arrgo

import (
	"fmt"
	"math"
)

//Empty 根据形状创建tensor.
//如果不指定shape, 则创建空的tensor.
//创建的tensor内部数据没有初始化。
func Empty(shape ...int) *Tensor {
	if len(shape) == 0 {
		return &Tensor{
			shape:   []int{},
			strides: []int{},
			data:    []float64{},
		}
	}

	var dNum = GetShapeSize(shape)
	var tShape = make([]int, len(shape))
	copy(tShape, shape)
	strides := make([]int, len(tShape)+1)
	strides[len(shape)] = 1
	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = strides[i+1] * tShape[i]
	}

	return &Tensor{
		shape:   tShape,
		strides: strides,
		data:    make([]float64, dNum),
	}
}

//EmptyLike 根据input Tensor的形状创建空Tensor.
//创建的tensor内部数据没有初始化。
func EmptyLike(input *Tensor) *Tensor {
	return Empty(input.shape...)
}

//Ones 根据指定形状创建内部数据泉为1.0的tensor.
func Ones(shape ...int) *Tensor {
	tenser := Empty(shape...)
	for i := range tenser.data {
		tenser.data[i] = 1.0
	}
	return tenser
}

//OnesLike 根据输入Tensor创建内部数据泉为1.0的tensor.
func OnesLike(input *Tensor) *Tensor {
	tenser := Empty(input.shape...)
	for i := range tenser.data {
		tenser.data[i] = 1.0
	}
	return tenser
}

//Zeros 根据指定形状创建内部数据泉为0.0的tensor.
func Zeros(shape ...int) *Tensor {
	tTensor := Empty(shape...)
	for i := range tTensor.data {
		tTensor.data[i] = 0.0
	}
	return tTensor
}

//ZerosLike 根据输入Tensor形状创建内部数据泉为0.0的tensor.
func ZerosLike(input *Tensor) *Tensor {
	tTensor := Empty(input.shape...)
	for i := range tTensor.data {
		tTensor.data[i] = 0.0
	}
	return tTensor
}

//Fill 根据指定形状创建内部数据泉为value的tensor.
func Fill(value float64, shape ...int) *Tensor {
	tTensor := Empty(shape...)
	for i := range tTensor.data {
		tTensor.data[i] = value
	}
	return tTensor
}

//Full 根据指定形状创建内部数据泉为value的tensor.
func Full(fillValue float64, shape ...int) *Tensor {
	return Fill(fillValue, shape...)
}

//FullLike 根据input Tensor的形状创建内部数据泉为value的tensor.
func FullLike(fillValue float64, input *Tensor) *Tensor {
	return Fill(fillValue, input.shape...)
}

//NewTensor 根据data和shape创建tensor.
//如果data个数和shape指定个数不同，抛出异常。
func NewTensor(data []float64, shape ...int) *Tensor {
	if len(data) == 0 {
		return Empty()
	}
	if int(len(data)) != GetShapeSize(shape) {
		panic(SHAPE_ERROR)
	}

	tensor := Empty(shape...)
	tData := make([]float64, len(data))
	copy(tData, data)
	tensor.data = tData
	return tensor
}

//Arange 通过指定起始、终止和步进量来创建一维Tensor。
// 输入参数： vals，可以有三种情况，详见下面描述。
// 情况1：Arange(stop): 以0开始的序列，创建Array [0, 0+(-)1, ..., stop)，不包括stop，stop符号决定升降序。
// 情况2：Arange(start, stop):创建Array [start, start +(-)1, ..., stop)，如果start小于start则递增，否则递减。
// 情况3：Arange(start, stop, step):创建Array [start, start + step, ..., stop)，step符号决定升降序。
// 输入参数多于三个的都会被忽略。
// 输出序列为“整型数”序列。
func Arange(vals ...int) *Tensor {
	var start, stop, step int = 0, 0, 1

	switch len(vals) {
	case 0:
		fmt.Println("range function should have range")
		panic(PARAMETER_ERROR)
	case 1:
		if vals[0] <= 0 {
			step = -1
			stop = vals[0] + 1
		} else {
			stop = vals[0] - 1
		}
	case 2:
		if vals[1] < vals[0] {
			step = -1
			stop = vals[1] + 1
		} else {
			stop = vals[1] - 1
		}
		start = vals[0]
	default:
		if vals[1] < vals[0] {
			if vals[2] >= 0 {
				fmt.Println("increment should be negative.")
				panic(PARAMETER_ERROR)
			}
			stop = vals[1] + 1
		} else {
			if vals[2] <= 0 {
				fmt.Println("increment should be positive.")
				panic(PARAMETER_ERROR)
			}
			stop = vals[1] - 1
		}
		start, step = vals[0], vals[2]
	}

	a := Empty(int(math.Abs(float64((stop-start)/step))) + 1)
	for i, v := 0, start; i < len(a.data); i, v = i+1, v+step {
		a.data[i] = float64(v)
	}
	return a
}

//Linspace 根据[start, stop]指定的区间，创建包含num个元素的一维数组。
func Linspace(start, stop float64, num int) *Tensor {
	var data = make([]float64, num)
	var startF, stopF = start, stop
	if startF <= stopF {
		var step = (stopF - startF) / (float64(num - 1.0))
		for i := range data {
			data[i] = startF + float64(i)*step
		}
		return NewTensor(data, num)
	} else {
		var step = (startF - stopF) / (float64(num - 1.0))
		for i := range data {
			data[i] = startF - float64(i)*step
		}
		return NewTensor(data, num)
	}
}

//Logspace 根据[10^start, 10^stop]指定的区间，创建包含num个元素的一维数组。
func Logspace(start, stop float64, num int) *Tensor {
	var data = make([]float64, num)
	var startF, stopF = start, stop
	if startF <= stopF {
		var step = (stopF - startF) / (float64(num - 1.0))
		for i := range data {
			data[i] = startF + float64(i)*step
		}
		return NewTensor(data, num)
	} else {
		var step = (startF - stopF) / (float64(num - 1.0))
		for i := range data {
			data[i] = math.Pow(10, startF-float64(i)*step)
		}
		return NewTensor(data, num)
	}
}

//Eye 返回2维Tensor，对角线为1.0，其余为0.0。
func Eye(n int) *Tensor {
	tensor := Zeros(n, n)
	for i := 0; i < n; i++ {
		tensor.data[i*(1+n)] = 1.0
	}
	return tensor
}

//Copy 复制Tensor
func Copy(input *Tensor) *Tensor {
	nShape := make([]int, len(input.shape))
	copy(nShape, input.shape)
	nData := make([]float64, len(input.data))
	copy(nData, input.data)
	return NewTensor(nData, nShape...)
}

func (t *Tensor) Copy() *Tensor {
	return Copy(t)
}
