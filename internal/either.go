package internal

func MapLeft[NL any, L any, R any](e Either[L, R], f func(v L) NL) Either[NL, R] {
	return NewLeft[NL, R](f(e.Left.value))
}

func MapRight[NR any, L any, R any](e Either[L, R], f func(v R) NR) Either[L, NR] {
	return NewRight[L, NR](f(e.Right.value))
}

type Left[L any] struct {
	value L
}

type Right[R any] struct {
	value R
}

type Either[L any, R any] struct {
	Left   Left[L]
	Right  Right[R]
	IsLeft bool
}

func NewLeft[L any, R any](left L) Either[L, R] {
	return Either[L, R]{Left: Left[L]{value: left}, IsLeft: true}
}

func NewRight[L any, R any](right R) Either[L, R] {
	return Either[L, R]{Right: Right[R]{value: right}}
}
