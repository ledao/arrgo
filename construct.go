package arrgo

func Empty(shape ...int64) *Tensor {
	if len(shape) == 0 {
		return &Tensor{
			shape:   []int64{},
			strides: []int64{},
			data:    []float64{},
		}

	} else {
		var d_num = GetShapeNum(shape)
		var t_shape = make([]int64, len(shape))
		copy(t_shape, shape)
		strides := make([]int64, len(t_shape)+1)
		strides[len(shape)] = 1
		for i := len(shape) - 1; i >= 0; i-- {
			strides[i] = strides[i+1] * t_shape[i]
		}

		return &Tensor{
			shape:   t_shape,
			strides: strides,
			data:    make([]float64, d_num),
		}
	}
}

func Ones(shape ...int64) *Tensor {
	t_dtenser := Empty(shape...)
	for i := range t_dtenser.data {
		t_dtenser.data[i] = 1.0
	}
	return t_dtenser
}

func New(data []float64, shape ...int64) *Tensor {

	return nil
}
