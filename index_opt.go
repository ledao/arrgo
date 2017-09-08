package arrgo

type Range struct {
    Start, Stop int
}

func (a *Arrf) Index(ranges ...Range) *Arrf {
    var ndim = len(a.shape)
    totalRanges := make([]Range, ndim)
    copy(totalRanges, ranges)
    if len(ranges) < ndim {
        for i := len(ranges); i < ndim; i++ {
            totalRanges[i] = Range{Start:0, Stop:a.shape[i]}
        }
    }

    shape := make([]int, ndim)
    for i := range shape {
        shape[i] = totalRanges[i].Stop - totalRanges[i].Start
    }

	b := Zeros(shape...)

    totalCount := 1
    for i := 0; i < ndim; i++ {
        totalCount *= shape[i]
    }

    counterSrc := make([]int, ndim)
    counterDst := make([]int, ndim)
    for i := range counterSrc {
        counterSrc[i] = totalRanges[i].Start
        counterDst[i] = counterSrc[i] - totalRanges[i].Start
    }

    for index := 0; index < totalCount; index++  {
        var v = a.At(counterSrc...)
        b.Set(v, counterDst...)
        counterSrc[ndim-1]++
        counterDst[ndim-1] = counterSrc[ndim-1] - totalRanges[ndim-1].Start
        var j = ndim -1
        for{
            if j >  0 && counterSrc[j] == totalRanges[j].Stop {
                counterSrc[j-1] ++
                counterSrc[j] = totalRanges[j].Start
                counterDst[j-1] = counterSrc[j-1] - totalRanges[j-1].Start
                counterDst[j] = counterSrc[j] - totalRanges[j].Start
                j--
            } else {
                break
            }
        }
    }

    return b
}
