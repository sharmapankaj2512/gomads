package internal

import "errors"

type Success[S any] struct {
	value S
}

type Failure struct {
	value error
}

type Try[S any] struct {
	success   Success[S]
	err       Failure
	isSuccess bool
}

func (n Try[S]) Value() S {
	return n.success.value
}

func (n Try[S]) Error() string {
	return n.err.value.Error()
}

func (n Try[S]) IsSuccess() bool {
	return n.isSuccess
}

func SuccessFor[S any](value S) Try[S] {
	return Try[S]{success: Success[S]{value: value}, isSuccess: true}
}

func ErrorAs[S any](message string) Try[S] {
	return Try[S]{err: Failure{value: errors.New(message)}, isSuccess: false}
}

func Map[S any, NS any](try Try[S], f func(v S) NS) Try[NS] {
	return SuccessFor(f(try.success.value))
}

func MapError[S any](err Try[S], f func(v error) error) Try[S] {
	return ErrorAs[S](f(err.err.value).Error())
}
