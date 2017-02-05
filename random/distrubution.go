
package random

import (
    "github.com/ledao/arrgo"
    "time"
    "math/rand"
)

var(
    r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Seed(seed int64) {
    r.Seed(seed)
}

//Return a random matrix with data from the "standard normal" distribution.
//
//`randn` generates a matrix filled with random floats sampled from a
//univariate "normal" (Gaussian) distribution of mean 0 and variance 1.
//
//Parameters
//----------
//\\*args : Arguments
//Shape of the output.
//If given as N integers, each integer specifies the size of one
//dimension. If given as a tuple, this tuple gives the complete shape.
//
//Returns
//-------
//Z : matrix of floats
//A matrix of floating-point samples drawn from the standard normal
//distribution.
func Randn(shape ...int) *arrgo.Arrf {
    a := arrgo.Empty(shape...)
    for i := range a.Values() {
        a.Values()[i] = r.NormFloat64()
    }

    return a
}
