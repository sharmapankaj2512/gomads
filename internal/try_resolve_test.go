package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveSuccess(t *testing.T) {
	Resolve(SuccessFor(1))(func(x int) {
		assert.Equal(t, 1, x)
	})
}

func Resolve[S any](try Try[S]) func(f func(try S)) {
	return func(f func(value S)) {
		f(try.success.value)
	}
}
