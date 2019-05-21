package arrgo

type Tensor struct {
	shape   []int64
	strides []int64
	data    []float64
}
