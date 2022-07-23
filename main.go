package main

import (
	"errors"
	"fmt"
)

type userError struct{}

func (e *userError) Error() string {
	return "Пользовательская ошибка"
}

type errorWithCause struct {
	msg   string
	cause error
}

func (e *errorWithCause) Error() string {
	return e.msg
}

func (e *errorWithCause) Cause() error {
	return e.cause
}

func (e *errorWithCause) Unwrap() error {
	return e.cause
}

func NewErrorWithCause(msg string, err error) error {
	return &errorWithCause{cause: err, msg: msg}
}

var userErr = &userError{}

func getWrappedError() error {
	return NewErrorWithCause("Обёртка", userErr)
}

func main() {
	err := getWrappedError()
	if errors.Is(err, userErr) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Это не пользовательская ошибка")
	}
}
