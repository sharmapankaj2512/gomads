package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformSuccess(t *testing.T) {
	success := SuccessFor(1)
	transformed := Map(success, func(v int) int { return v + 1 })

	assert.Equal(t, 2, transformed.success.value)
}

func TestTransformError(t *testing.T) {
	err := ErrorAs[int]("something went wrong")
	transformed := MapError(err, func(v error) error { return errors.New("failed") })

	assert.Equal(t, "failed", transformed.err.value.Error())
}

type Success[S any] struct {
	value S
}

type Failure struct {
	value error
}

type Try[S any] struct {
	success Success[S]
	err     Failure
}

func SuccessFor[S any](value S) Try[S] {
	return Try[S]{success: Success[S]{value: value}}
}

func ErrorAs[S any](message string) Try[S] {
	return Try[S]{err: Failure{value: errors.New(message)}}
}

func Map[S any, NS any](try Try[S], f func(v S) NS) Try[NS] {
	return SuccessFor(f(try.success.value))
}

func MapError[S any](err Try[S], f func(v error) error) Try[S] {
	return ErrorAs[S](f(err.err.value).Error())
}
