package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 1)
	chan2 <- 2

	// The program either prints 1 or 2 and terminates
	select {
	case chan1Val := <-chan1:
		fmt.Println(chan1Val)
	case chan2Val := <-chan2:
		fmt.Println(chan2Val)
	}
}
