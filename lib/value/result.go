package value

import (
	"log"
	"main/lib"
)

type Result[T any] struct {
	Value T
	Error error
}

func (result *Result[T]) Ok() bool {
	return nil == result.Error
}

func Wrap[T any](value T, err error) Result[T] {
	if nil != err {
		lib.Notifier.SendError(err)
		return Result[T]{value, err}
	}

	return Result[T]{
		Value: value,
		Error: err,
	}
}

func WrapFatal[T any](value T, err error) Result[T] {
	if nil != err {
		log.Fatal(err)
	}

	return Result[T]{
		Value: value,
		Error: err,
	}
}

func WrapNothing(err error) Result[int] {
	if nil == err {
		return Result[int]{Value: 0}
	}
	lib.Notifier.SendError(err)
	return Result[int]{1, err}

}

func WrapNothingFatal(err error) Result[int] {
	if nil != err {
		log.Fatal(err)
	}
	return Result[int]{Value: 0}
}
