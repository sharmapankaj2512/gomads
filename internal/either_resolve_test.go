package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveWhenBothAreLeft(t *testing.T) {
	left1 := NewLeft[int, error](1)
	left2 := NewLeft[string, error]("1")

	ResolveLeft(left1, left2, func(v1 int, v2 string) {
		assert.Equal(t, 1, v1)
		assert.Equal(t, "1", v2)
	})
}

func TestDoNotResolveWhenEitherIsRight(t *testing.T) {
	left1 := NewLeft[int, error](1)
	left2 := NewRight[string, string]("1")

	ResolveLeft(left1, left2, func(v1 int, v2 string) {
		assert.Fail(t, "should not be invoked")
	})
}

func ResolveLeft[L any, R any, LL any, RR any](e1 Either[L, R], e2 Either[LL, RR], f func(v1 L, v2 LL)) {
	if e1.IsLeft && e2.IsLeft {
		f(e1.Left.value, e2.Left.value)
	}
}
