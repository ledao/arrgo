package arrgo

import (
    "github.com/ledao/arrgo/internal"
)

func (a *Arrf) AddC(b float64) *Arrf {
    asm.AddC(b, a.data)
    return a
}
