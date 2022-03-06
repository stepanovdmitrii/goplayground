package main

type Interface1 interface {
	Method() string
}

type Interface2 interface {
	Method() string
}

type Interface3 interface {
	Interface1
	Interface2
}

func main() {
	var i Interface3
	i.Method()
}
