package main

import "fmt"

func appendInt64(sl []int64) {
	sl = append(sl, 1000)
	fmt.Printf("%v\n", sl)
}

func main() {
	sl := make([]int64, 0, 10)
	for i := 0; i < 9; i++ {
		sl = append(sl, int64(i))
	}
	appendInt64(sl)
	fmt.Printf("%v\n", sl)
}
