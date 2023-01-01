package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformSuccess(t *testing.T) {
	success := SuccessFor(1)
	transformed := Map(success, func(v int) int { return v + 1 })

	assert.Equal(t, 2, transformed.success)
}

type Try[S any] struct {
	success S
	err     error
}

func SuccessFor[S any](value S) Try[S] {
	return Try[S]{success: value}
}

func Map[S any, NS any](try Try[S], f func(v S) NS) Try[NS] {
	return SuccessFor(f(try.success))
}
