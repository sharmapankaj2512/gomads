package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveLeftOne(t *testing.T) {
	Pipe(MapLeft(NewLeft[int, error](1), func(v int) int { return v + 1 }),
		ResolveLeftOne[int, error](func(v1 int) {
			assert.Equal(t, 2, v1)
		}))
}
