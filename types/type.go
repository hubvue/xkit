package xkit

type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type ComplexNumber interface {
	~complex64 | ~complex128
}

type Number interface {
	RealNumber | ComplexNumber
}
