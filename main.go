package main

import "fmt"

type Inner struct {

}

func (s *Inner) Hello() {
	fmt.Println("inner")
}


type Outer struct {
	Inner
}

func (s *Outer) Hello() {
	fmt.Println("outer")
}



func main() {
	o := &Outer{}
	o.Hello()
	o.Inner.Hello()
}
