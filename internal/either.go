package internal

type Left[L any] struct {
	value L
}

type Right[R any] struct {
	value R
}

type Either[L any, R any] struct {
	Left    Left[L]
	Right   Right[R]
	IsLeft  bool
	IsRight bool
}

func MapLeft[NL any, L any, R any](e Either[L, R], f func(v L) NL) Either[NL, R] {
	return NewLeft[NL, R](f(e.Left.value))
}

func MapRight[NR any, L any, R any](e Either[L, R], f func(v R) NR) Either[L, NR] {
	return NewRight[L, NR](f(e.Right.value))
}

func NewLeft[L any, R any](left L) Either[L, R] {
	return Either[L, R]{Left: Left[L]{value: left}, IsLeft: true, IsRight: false}
}

func NewRight[L any, R any](right R) Either[L, R] {
	return Either[L, R]{Right: Right[R]{value: right}, IsLeft: false, IsRight: true}
}

func Pipe[L any, R any](left Either[L, R], f func(newLeft Either[L, R])) {
	f(left)
}

func ResolveLeftOne[L any, R any](e1 Either[L, R], f func(v1 L)) {
	if e1.IsLeft {
		f(e1.Left.value)
	}
}

func ResolveLeftTwo[L any, R any, LL any, RR any](e1 Either[L, R], e2 Either[LL, RR], f func(v1 L, v2 LL)) {
	if e1.IsLeft && e2.IsLeft {
		f(e1.Left.value, e2.Left.value)
	}
}

func ResolveRightTwo[L any, R any, LL any, RR any](e1 Either[L, R], e2 Either[LL, RR], f func(v1 R, v2 RR)) {
	if e1.IsRight && e2.IsRight {
		f(e1.Right.value, e2.Right.value)
	}
}
