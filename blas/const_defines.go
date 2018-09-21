package blas

type CBLAS_ORDER int64
const CblasRowMajor: CBLAS_ORDER = 101
const CblasColMajor: CBLAS_ORDER = 102

type CBLAS_TRANSPOSE int64
const CblasNoTrans :CBLAS_TRANSPOSE = 111
const CblasTrans :CBLAS_TRANSPOSE = 112
const CblasConjTrans: CBLAS_TRANSPOSE = 113
const CblasConjNoTrans: CBLAS_TRANSPOSE = 114

type CBLAS_UPLO int64
const CblasUpper: CBLAS_UPLO = 121
const CblasLower: CBLAS_UPLO = 122

type CBLAS_DIAG int64
const CblasNonUnit: CBLAS_DIAG = 131
const CblasUnit: CBLAS_DIAG = 132

type CBLAS_SIDE int64
const CblasLeft: CBLAS_SIDE = 141
const CblasRight: CBLAS_SIDE = 142
