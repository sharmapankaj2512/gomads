package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	left := NewLeft[int, error](1)
	newLeft := MapLeft(left, func(v int) int { return v + 1 })
	ResolveLeftOne(newLeft, func(v int) {
		assert.Equal(t, 2, v)
	})
}
