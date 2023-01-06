package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveSuccess(t *testing.T) {
	Resolve(SuccessFor(1))(
		func(x int) {
			assert.Equal(t, 1, x)
		})
}

func TestResolveError(t *testing.T) {
	ResolveError(ErrorAs[int]("something went wrong"))(
		func(x error) {
			assert.Equal(t, "something went wrong", x.Error())
		})
}

func Resolve[S any](try Try[S]) func(f func(try S)) {
	return func(f func(value S)) {
		f(try.success.value)
	}
}

func ResolveError[S any](try Try[S]) func(f func(value error)) {
	return func(f func(value error)) {
		f(try.err.value)
	}
}
