package main

import "fmt"

type ErrType struct {
}

func (e *ErrType) Error() string {
	return "error msg"
}

func createErr() error {
	var errType *ErrType = nil
	return errType
}

func main() {
	errType := createErr()
	fmt.Printf("%t", errType == nil)
}
