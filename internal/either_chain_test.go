package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveLeftOne(t *testing.T) {
	left := NewLeft[int, error](1)
	addOne := func(v int) int { return v + 1 }
	assertIsTwo := func(v1 int) { assert.Equal(t, 2, v1) }

	Pipe(
		MapLeft(left, addOne),
		ResolveLeftOne[int, error](assertIsTwo))
}
