package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapAndResolve(t *testing.T) {
	successFor := SuccessFor(1)
	Resolve(Map(successFor, func(v int) int { return v + 1 }))(
		func(v int) {
			assert.Equal(t, 2, v)
		})
}
