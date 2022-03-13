package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- 100
		time.Sleep(time.Second * 5)
		close(ch2)
	}()
	close(ch1)
	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println("from ch1")
		case _, ok := <-ch2:
			if !ok {
				ch2  = nil
				break
			}
			fmt.Println("from ch2")
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
