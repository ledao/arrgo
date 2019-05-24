package arrgo

//Size tensor内部元素个数
func (t *Tensor) Size() int {
	return int(len(t.data))
}

//Numel tensor内部元素个数
func (t *Tensor) Numel() int {
	return t.Size()
}

//NDims 获取Tensor维度个数
func (t *Tensor) NDims() int {
	return len(t.shape)
}
