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
		5, 10)).Any() {
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
		2, 10)).Any() {
		t.Error(`Expected 
				[[100, 105, 110, 115, 120, 125, 130, 135, 140, 145],
 				[350, 355, 360, 365, 370, 375, 380, 385, 390, 395]], got `,
			arr.Sum(1))
	}

	if arr.Sum(2).NotEqual(Array(
		[]float64{
			45, 145, 245, 345, 445,
			545, 645, 745, 845, 945},
		2, 5)).Any() {
		t.Error(`Expected 
			[[ 45, 145, 245, 345, 445],
	 		[545, 645, 745, 845, 945]]
		, got 
		`, arr.Sum(2))
	}

	if arr.Sum(0, 1).NotEqual(Array(
		[]float64{
			450, 460, 470, 480, 490, 500, 510, 520, 530, 540},
		10)).Any() {
		t.Error(`Expected 
			[450, 460, 470, 480, 490, 500, 510, 520, 530, 540]
			, got`,
			arr.Sum(0, 1))
	}

	if arr.Sum(0, 2).NotEqual(Array([]float64{590, 790, 990, 1190, 1390}, 5)).Any() {
		t.Error(`Expected [ 590,  790,  990, 1190, 1390], got `, arr.Sum(0, 2))
	}

	if arr.Sum(1, 2).NotEqual(Array([]float64{1225, 3725})).Any() {
		t.Error("Expected [1225, 3725], got ", arr.Sum(1, 2))
	}

	if arr.Sum(0, 1, 2).NotEqual(Array([]float64{4950})).Any() {
		t.Error("expected [4950], got ", arr.Sum(0, 1, 2))
	}

	if arr.Sum().NotEqual(Array([]float64{4950})).Any() {
		t.Error("expected [4950], got ", arr.Sum())
	}
}
