package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	left := NewLeft[int, error](1)
	Pipe(MapLeft(left, func(v int) int { return v + 1 }),
		func(newLeft Either[int, error]) {
			ResolveLeftOne(newLeft, func(v int) {
				assert.Equal(t, 2, v)
			})
		})
}

func Pipe[L any, R any](left Either[L, R], f func(newLeft Either[L, R])) {
	f(left)
}
