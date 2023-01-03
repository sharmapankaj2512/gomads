package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformSuccess(t *testing.T) {
	success := SuccessFor(1)
	transformed := Map(success, func(v int) int { return v + 1 })

	assert.Equal(t, 2, transformed.Value())
}

func TestTransformError(t *testing.T) {
	err := ErrorAs[int]("something went wrong")
	transformed := MapError(err, func(v error) error { return errors.New("failed") })

	assert.Equal(t, "failed", transformed.Error())
}
