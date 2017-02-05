package arrgo

import (
    "math"
)

func ReverseIntSlice(slice []int) []int {
    s := make([]int, len(slice))
    copy(s, slice)
    for i, j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

func ProductIntSlice(slice []int) int {
    var prod = 1
    for _, v := range slice {
        prod *= v
    }
    return prod
}


func Roundf(val float64, places int) float64 {
    var t float64
    f := math.Pow10(places)
    x := val * f
    if math.IsInf(x, 0) || math.IsNaN(x) {
        return val
    }
    if x >= 0.0 {
        t = math.Ceil(x)
        if (t - x) > 0.50000000001 {
            t -= 1.0
        }
    } else {
        t = math.Ceil(-x)
        if (t + x) > 0.50000000001 {
            t -= 1.0
        }
        t = -t
    }
    x = t / f

    if !math.IsInf(x, 0) {
        return x
    }

    return t
}