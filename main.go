package main

import "fmt"

func main() {
	var x = 3
	switch {
	case x%2 == 0:
		fmt.Println("2")
		fallthrough
	case x%3 == 0:
		fmt.Println("3")
	case x%4 == 0:
		fmt.Println("4")
		fallthrough
	case x%5 == 0:
		fmt.Println("5")
		fallthrough
	default:
		fmt.Println("none")
	}
}
