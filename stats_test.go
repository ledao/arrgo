package arrgo

import "testing"

func TestSum(t *testing.T) {
	var arr = Arange(100).Reshape(2, 5, 10)
	if arr.Sum(0).NotEqual(Array(
		[]float64{
			50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
			70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
			90, 92, 94, 96, 98, 100, 102, 104, 106, 108,
			110, 112, 114, 116, 118, 120, 122, 124, 126, 128,
			130, 132, 134, 136, 138, 140, 142, 144, 146, 148},
		5, 10)).AnyTrue() {
		t.Error(`Expected [[ 50,  52,  54,  56,  58,  60,  62,  64,  66,  68],
							 [ 70,  72,  74,  76,  78,  80,  82,  84,  86,  88],
							 [ 90,  92,  94,  96,  98, 100, 102, 104, 106, 108],
							 [110, 112, 114, 116, 118, 120, 122, 124, 126, 128],
							 [130, 132, 134, 136, 138, 140, 142, 144, 146, 148]], got `,
			arr.Sum(0))
	}

	if arr.Sum(1).NotEqual(Array(
		[]float64{
			100, 105, 110, 115, 120, 125, 130, 135, 140, 145,
			350, 355, 360, 365, 370, 375, 380, 385, 390, 395},
		2, 10)).AnyTrue() {
		t.Error(`Expected 
				[[100, 105, 110, 115, 120, 125, 130, 135, 140, 145],
 				[350, 355, 360, 365, 370, 375, 380, 385, 390, 395]], got `,
			arr.Sum(1))
	}

	if arr.Sum(2).NotEqual(Array(
		[]float64{
			45, 145, 245, 345, 445,
			545, 645, 745, 845, 945},
		2, 5)).AnyTrue() {
		t.Error(`Expected 
			[[ 45, 145, 245, 345, 445],
	 		[545, 645, 745, 845, 945]]
		, got 
		`, arr.Sum(2))
	}

	if arr.Sum(0, 1).NotEqual(Array(
		[]float64{
			450, 460, 470, 480, 490, 500, 510, 520, 530, 540},
		10)).AnyTrue() {
		t.Error(`Expected 
			[450, 460, 470, 480, 490, 500, 510, 520, 530, 540]
			, got`,
			arr.Sum(0, 1))
	}

	if arr.Sum(0, 2).NotEqual(Array([]float64{590, 790, 990, 1190, 1390}, 5)).AnyTrue() {
		t.Error(`Expected [ 590,  790,  990, 1190, 1390], got `, arr.Sum(0, 2))
	}

	if arr.Sum(1, 2).NotEqual(Array([]float64{1225, 3725})).AnyTrue() {
		t.Error("Expected [1225, 3725], got ", arr.Sum(1, 2))
	}

	if arr.Sum(0, 1, 2).NotEqual(Array([]float64{4950})).AnyTrue() {
		t.Error("expected [4950], got ", arr.Sum(0, 1, 2))
	}

	if arr.Sum().NotEqual(Array([]float64{4950})).AnyTrue() {
		t.Error("expected [4950], got ", arr.Sum())
	}
}


func TestArgMax(t *testing.T) {
	arr := Array([]float64{17, 10, 22,  3,  2,  7, 15,  9, 23,  4, 14, 18,  5,  8,  0, 12,  1,
			       19, 20, 11,  6, 16, 21, 13}, 2,3,4)

	if arr.ArgMax(0).NotEqual(Array([]float64{0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0}, 3, 4)).AnyTrue() {
		t.Error(`Expected
		[[0, 0, 0, 1],
		[0, 1, 1, 1],
		[0, 1, 1, 0]], got `, arr.ArgMax(0))
	}

	if arr.ArgMax(1).NotEqual(Array([]float64{2, 0, 0, 2, 2, 1, 2, 2}, 2, 4)).AnyTrue() {
		t.Error(`Expected
		[[2, 0, 0, 2],
       		[2, 1, 2, 2]], got `, arr.ArgMax(1))
	}

	if arr.ArgMax(2).NotEqual(Array([]float64{2, 2, 0, 3, 2, 2}, 2, 3)).AnyTrue() {
		t.Error(`Expected
		[[2, 2, 0],
       		[3, 2, 2]], got `, arr.ArgMax(2))
	}
}

func TestArgMin(t *testing.T) {
	arr := Array([]float64{17, 10, 22,  3,  2,  7, 15,  9, 23,  4, 14, 18,  5,  8,  0, 12,  1,
			       19, 20, 11,  6, 16, 21, 13}, 2,3,4)

	if arr.ArgMin(0).NotEqual(Array([]float64{1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1}, 3, 4)).AnyTrue() {
		t.Error(`Expected
		[[1, 1, 1, 0],
		[1, 0, 0, 0],
		[1, 0, 0, 1]], got `, arr.ArgMin(0))
	}

	if arr.ArgMin(1).NotEqual(Array([]float64{1, 2, 2, 0, 1, 0, 0, 1}, 2, 4)).AnyTrue() {
		t.Error(`Expected
		[[1, 2, 2, 0],
       		[1, 0, 0, 1]], got `, arr.ArgMin(1))
	}

	if arr.ArgMin(2).NotEqual(Array([]float64{3, 0, 1, 2, 0, 0}, 2, 3)).AnyTrue() {
		t.Error(`Expected
		[[3, 0, 1],
       		[2, 0, 0]], got `, arr.ArgMin(2))
	}
}