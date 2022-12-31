package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveLeft(t *testing.T) {
	left := NewLeft[int, error](1)

	ResolveLeftOne(left, func(v int) {
		assert.Equal(t, 1, v)
	})
}

func TestResolveWhenBothAreLeft(t *testing.T) {
	left1 := NewLeft[int, error](1)
	left2 := NewLeft[string, error]("1")

	ResolveLeft(left1, left2, func(v1 int, v2 string) {
		assert.Equal(t, 1, v1)
		assert.Equal(t, "1", v2)
	})
}

func TestDoNotResolveWhenEitherIsRight(t *testing.T) {
	left := NewLeft[int, error](1)
	right := NewRight[string, string]("1")

	ResolveLeft(left, right, func(v1 int, v2 string) {
		assert.Fail(t, "should not be invoked")
	})
}

func TestResolveWhenBothAreRight(t *testing.T) {
	right1 := NewRight[string, int](1)
	right2 := NewRight[string, string]("1")

	ResolveRight(right1, right2, func(v1 int, v2 string) {
		assert.Equal(t, 1, v1)
		assert.Equal(t, "1", v2)
	})
}

func TestDoNotResolveWhenEitherIsLeft(t *testing.T) {
	left := NewLeft[int, error](1)
	right := NewRight[string, string]("1")

	ResolveRight(left, right, func(v1 error, v2 string) {
		assert.Fail(t, "should not be invoked")
	})
}

func ResolveRight[L any, R any, LL any, RR any](e1 Either[L, R], e2 Either[LL, RR], f func(v1 R, v2 RR)) {
	if e1.IsRight && e2.IsRight {
		f(e1.Right.value, e2.Right.value)
	}
}

func ResolveLeftOne[L any, R any](e1 Either[L, R], f func(v1 L)) {
	if e1.IsLeft {
		f(e1.Left.value)
	}
}

func ResolveLeft[L any, R any, LL any, RR any](e1 Either[L, R], e2 Either[LL, RR], f func(v1 L, v2 LL)) {
	if e1.IsLeft && e2.IsLeft {
		f(e1.Left.value, e2.Left.value)
	}
}
