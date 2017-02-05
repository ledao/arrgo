package arrgo

import "fmt"

var (
    INDEX_ERROR error = fmt.Errorf("INDEX ERROR")
    SHAPE_ERROR error = fmt.Errorf("SHAPE ERROR")
    DIMENTION_ERROR error = fmt.Errorf("DIMENTION ERROR")
    TYPE_ERROR error = fmt.Errorf("DATA TYPE ERROR")

    UNIMPLEMENT_ERROR error = fmt.Errorf("UNIMPLEMENT ERROR")
)
