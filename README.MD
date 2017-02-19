### type Arrf
    func Abs(b *Arrf) *Arrf
    func Acos(b *Arrf) *Arrf
    func Acosh(b *Arrf) *Arrf
    func Add(a, b *Arrf) *Arrf
    func Arange(vals ...float64) (a *Arrf)
    func ArgMax(a *Arrf, axis ...int) *Arrf
    func ArgMin(a *Arrf, axis ...int) *Arrf
    func Array(data []float64, shape ...int) (a *Arrf)
    func Asin(b *Arrf) *Arrf
    func Asinh(b *Arrf) *Arrf
    func Atan(b *Arrf) *Arrf
    func Atanh(b *Arrf) *Arrf
    func Ceil(b *Arrf) *Arrf
    func CopySign(a, b *Arrf) *Arrf
    func Cos(b *Arrf) *Arrf
    func Cosh(b *Arrf) *Arrf
    func Div(a, b *Arrf) *Arrf
    func Empty(shape ...int) (a *Arrf)
    func EmptyLike(a *Arrf) *Arrf
    func Exp(b *Arrf) *Arrf
    func Eye(n int) *Arrf
    func Floor(b *Arrf) *Arrf
    func Full(fullValue float64, shape ...int) *Arrf
    func Identity(n int) *Arrf
    func Log(b *Arrf) *Arrf
    func Log10(b *Arrf) *Arrf
    func Log1p(b *Arrf) *Arrf
    func Log2(b *Arrf) *Arrf
    func Max(a *Arrf, axis ...int) *Arrf
    func Maximum(a, b *Arrf) *Arrf
    func Mean(a *Arrf, axis ...int) *Arrf
    func Min(a *Arrf, axis ...int) *Arrf
    func Minimum(a, b *Arrf) *Arrf
    func Mod(a, b *Arrf) *Arrf
    func Modf(b *Arrf) (*Arrf, *Arrf)
    func Mul(a, b *Arrf) *Arrf
    func Ones(shape ...int) *Arrf
    func OnesLike(a *Arrf) *Arrf
    func Pow(a, b *Arrf) *Arrf
    func Round(b *Arrf, places int) *Arrf
    func Sign(b *Arrf) *Arrf
    func Sin(b *Arrf) *Arrf
    func Sinh(b *Arrf) *Arrf
    func Sort(a *Arrf, axis ...int) *Arrf
    func Sqrt(b *Arrf) *Arrf
    func Square(b *Arrf) *Arrf
    func Std(a *Arrf, axis ...int) *Arrf
    func Sub(a, b *Arrf) *Arrf
    func Sum(a *Arrf, axis ...int) *Arrf
    func Tan(b *Arrf) *Arrf
    func Tanh(b *Arrf) *Arrf
    func Var(a *Arrf, axis ...int) *Arrf
    func Where(cond *Arrb, tv, fv interface{}) *Arrf
    func Zeros(shape ...int) *Arrf
    func ZerosLike(a *Arrf) *Arrf
    func (a *Arrf) Add(b *Arrf) *Arrf
    func (a *Arrf) AddC(b float64) *Arrf
    func (a *Arrf) ArgMax(axis ...int) *Arrf
    func (a *Arrf) ArgMin(axis ...int) *Arrf
    func (a *Arrf) At(index ...int) float64
    func (a *Arrf) Copy() *Arrf
    func (a *Arrf) Count(axis ...int) int
    func (a *Arrf) Div(b *Arrf) *Arrf
    func (a *Arrf) DivC(b float64) *Arrf
    func (a *Arrf) DotProd(b *Arrf) float64
    func (a *Arrf) Equal(b *Arrf) *Arrb
    func (a *Arrf) Flatten() *Arrf
    func (a *Arrf) Get(index ...int) float64
    func (a *Arrf) Greater(b *Arrf) *Arrb
    func (a *Arrf) GreaterEqual(b *Arrf) *Arrb
    func (a *Arrf) Index(ranges ...Range) *Arrf
    func (a *Arrf) Less(b *Arrf) *Arrb
    func (a *Arrf) LessEqual(b *Arrf) *Arrb
    func (a *Arrf) MatProd(b *Arrf) *Arrf
    func (a *Arrf) Max(axis ...int) *Arrf
    func (a *Arrf) Mean(axis ...int) *Arrf
    func (a *Arrf) Min(axis ...int) *Arrf
    func (a *Arrf) Mul(b *Arrf) *Arrf
    func (a *Arrf) MulC(b float64) *Arrf
    func (a *Arrf) Ndims() int
    func (a *Arrf) NotEqual(b *Arrf) *Arrb
    func (a *Arrf) Reshape(shape ...int) *Arrf
    func (a *Arrf) Set(val float64, index ...int) *Arrf
    func (a *Arrf) Sort(axis ...int) *Arrf
    func (a *Arrf) Std(axis ...int) *Arrf
    func (a *Arrf) String() (s string)
    func (a *Arrf) Sub(b *Arrf) *Arrf
    func (a *Arrf) SubC(b float64) *Arrf
    func (a *Arrf) Sum(axis ...int) *Arrf
    func (a *Arrf) Transpose(axes ...int) *Arrf
    func (a *Arrf) Values() []float64
    func (a *Arrf) Var(axis ...int) *Arrf

### type Arrb
    func ArrayB(data []bool, shape ...int) (a *Arrb)
    func EmptyB(shape ...int) (a *Arrb)
    func Equal(a, b *Arrf) *Arrb
    func FullB(value bool, shape ...int) *Arrb
    func Greater(a, b *Arrf) *Arrb
    func GreaterEqual(a, b *Arrf) *Arrb
    func IsFinit(b *Arrf) *Arrb
    func IsInf(b *Arrf) *Arrb
    func IsNaN(b *Arrf) *Arrb
    func Less(a, b *Arrf) *Arrb
    func LessEqual(a, b *Arrf) *Arrb
    func LogicalAnd(a, b *Arrb) *Arrb
    func LogicalNot(a *Arrb) *Arrb
    func LogicalOr(a, b *Arrb) *Arrb
    func NotEqual(a, b *Arrf) *Arrb
    func (ab *Arrb) All() bool
    func (ab *Arrb) Any() bool
    func (a *Arrb) LogicalAnd(b *Arrb) *Arrb
    func (a *Arrb) LogicalNot() *Arrb
    func (a *Arrb) LogicalOr(b *Arrb) *Arrb
    func (a *Arrb) String() (s string)
    func (a *Arrb) Sum() int

### Variables
    func ContainsFloat64(s []float64, e float64) bool
    func Hargmax(ln int, data []float64)
    func Hargmin(ln int, data []float64)
    func Hmax(ln int, data []float64)
    func Hmin(ln int, data []float64)
    func Hsort(ln int, data []float64)
    func ProductIntSlice(slice []int) int
    func ReverseIntSlice(slice []int) []int
    func Roundf(val float64, places int) float64
    func Vargmax(ln int, a []float64)
    func Vargmin(ln int, a []float64)
    func Vmax(a, b []float64)
    func Vmin(a, b []float64)
    func Vsort(ln int, a []float64)