package arrgo

import (
    "github.com/ledao/arrgo/internal"
)

func (a *Arrgo) AddC(b float64) *Arrgo {
    asm.AddC(b, a.data)
    return a
}
