---
title: "Go Math"
date: 2020-12-02T09:53:41+08:00
draft: false
---

Package math

### go math page 定义的常量

```go
E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622

Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339

Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
Log2E  = 1 / Ln2
Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
Log10E = 1 / Ln10


// Floating-point limit values
MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)

// Integer limit values.
MaxInt8   = 1<<7 - 1
MinInt8   = -1 << 7
MaxInt16  = 1<<15 - 1
MinInt16  = -1 << 15
MaxInt32  = 1<<31 - 1
MinInt32  = -1 << 31
MaxInt64  = 1<<63 - 1
MinInt64  = -1 << 63
MaxUint8  = 1<<8 - 1
MaxUint16 = 1<<16 - 1
MaxUint32 = 1<<32 - 1
MaxUint64 = 1<<64 - 1

fmt.Println(math.MaxInt8)
// 127
```


### go math page 定义的方法

```go
func Abs(x float64) float64
func Acos(x float64) float64
func Acosh(x float64) float64
func Asin(x float64) float64
func Asinh(x float64) float64
func Atan(x float64) float64
func Atan2(y, x float64) float64
func Atanh(x float64) float64
func Cbrt(x float64) float64
func Ceil(x float64) float64
func Copysign(x, y float64) float64
func Cos(x float64) float64
func Cosh(x float64) float64
func Dim(x, y float64) float64
func Erf(x float64) float64
func Erfc(x float64) float64
func Erfcinv(x float64) float64
func Erfinv(x float64) float64
func Exp(x float64) float64
func Exp2(x float64) float64
func Expm1(x float64) float64
func FMA(x, y, z float64) float64
func Float32bits(f float32) uint32
func Float32frombits(b uint32) float32
func Float64bits(f float64) uint64
func Float64frombits(b uint64) float64
func Floor(x float64) float64
func Frexp(f float64) (frac float64, exp int)
func Gamma(x float64) float64
func Hypot(p, q float64) float64
func Ilogb(x float64) int
func Inf(sign int) float64
func IsInf(f float64, sign int) bool
func IsNaN(f float64) (is bool)
func J0(x float64) float64
func J1(x float64) float64
func Jn(n int, x float64) float64
func Ldexp(frac float64, exp int) float64
func Lgamma(x float64) (lgamma float64, sign int)
func Log(x float64) float64
func Log10(x float64) float64
func Log1p(x float64) float64
func Log2(x float64) float64
func Logb(x float64) float64
func Max(x, y float64) float64
func Min(x, y float64) float64
func Mod(x, y float64) float64
func Modf(f float64) (int float64, frac float64)
func NaN() float64
func Nextafter(x, y float64) (r float64)
func Nextafter32(x, y float32) (r float32)
func Pow(x, y float64) float64
func Pow10(n int) float64
func Remainder(x, y float64) float64
func Round(x float64) float64
func RoundToEven(x float64) float64
func Signbit(x float64) bool
func Sin(x float64) float64
func Sincos(x float64) (sin, cos float64)
func Sinh(x float64) float64
func Sqrt(x float64) float64
func Tan(x float64) float64
func Tanh(x float64) float64
func Trunc(x float64) float64
func Y0(x float64) float64
func Y1(x float64) float64
func Yn(n int, x float64) float64

// 取绝对值
math.Abs(-2)


math.Acos(1)  // 0

math.Acosh(1) // 0.00

math.Asin(0) // 0.00

math.Asinh(0) // 0.00

math.Atan(0)
math.Atan2(0, 0)
math.Atanh(0)

math.Cbrt(8)
math.Ceil(1.49)
math.Copysign(3.2, -1)
math.Cos(math.Pi/2)
math.Cosh(0)

math.Dim(4, -2)  // 6
math.Dim(-4, 2)   // 0.00
Dim(+Inf, +Inf) = NaN
Dim(-Inf, -Inf) = NaN
Dim(x, NaN) = Dim(NaN, x) = NaN

Erf(+Inf) = 1
Erf(-Inf) = -1
Erf(NaN) = NaN

Erfc(+Inf) = 0
Erfc(-Inf) = 2
Erfc(NaN) = NaN

Erfcinv(0) = +Inf
Erfcinv(2) = -Inf
Erfcinv(x) = NaN if x < 0 or x > 2
Erfcinv(NaN) = NaN

Erfinv(1) = +Inf
Erfinv(-1) = -Inf
Erfinv(x) = NaN if x < -1 or x > 1
Erfinv(NaN) = NaN

Exp(+Inf) = +Inf
Exp(NaN) = NaN

math.Exp(1) // 2.72
math.Exp(2) // 7.39
math.Exp(-1) // 0.37

math.Exp2(1)
math.Exp2(-3)

Abs
Acos
Acosh
Asin
Asinh
Atan
Atan2
Atanh
Cbrt
Ceil
Copysign
Cos
Cosh
Dim
Exp
Exp2
Expm1
Floor
Log
Log10
Log2
Mod
Modf
Pow
Pow10
Round
RoundToEven
Sin
Sincos
Sinh
Sqrt
Tan
Tanh
Trunc
 ```


 ### 参考资料

 - [go math](https://golang.org/pkg/math/)
 - [Package cmplx](https://golang.org/pkg/math/cmplx/)