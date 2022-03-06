package main

import "fmt"

type TPtr struct {

}

func (t *TPtr) Hello() {
	fmt.Println("hello from TPtr")
}



type TValue struct {

}

func (t *TValue) Hello() {
	fmt.Println("hello from TValue")
}


func main() {
	pTPtr := &TPtr{}
	pTPtr.Hello()

	vTPtr := TPtr{}
	vTPtr.Hello()


	pTValue := &TValue{}
	pTValue.Hello()

	vTValue := TValue{}
	vTValue.Hello()
}
