package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveLeft(t *testing.T) {
	left := NewLeft[int, error](1)

	ResolveLeftOne[int, error](func(v int) {
		assert.Equal(t, 1, v)
	})(left)
}

func TestResolveWhenBothAreLeft(t *testing.T) {
	left1 := NewLeft[int, error](1)
	left2 := NewLeft[string, error]("1")

	ResolveLeftTwo(left1, left2, func(v1 int, v2 string) {
		assert.Equal(t, 1, v1)
		assert.Equal(t, "1", v2)
	})
}

func TestDoNotResolveWhenEitherIsRight(t *testing.T) {
	left := NewLeft[int, error](1)
	right := NewRight[string, string]("1")

	ResolveLeftTwo(left, right, func(v1 int, v2 string) {
		assert.Fail(t, "should not be invoked")
	})
}

func TestResolveWhenBothAreRight(t *testing.T) {
	right1 := NewRight[string, int](1)
	right2 := NewRight[string, string]("1")

	ResolveRightTwo(right1, right2, func(v1 int, v2 string) {
		assert.Equal(t, 1, v1)
		assert.Equal(t, "1", v2)
	})
}

func TestDoNotResolveWhenEitherIsLeft(t *testing.T) {
	left := NewLeft[int, error](1)
	right := NewRight[string, string]("1")

	ResolveRightTwo(left, right, func(v1 error, v2 string) {
		assert.Fail(t, "should not be invoked")
	})
}
