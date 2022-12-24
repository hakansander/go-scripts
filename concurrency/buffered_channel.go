package main

import "fmt"

func main() {
	bufferedChannel := make(chan int, 5)
	done := make(chan bool)

	go func(bufferedChannel chan int, done chan bool) {
		//when the buf channel closed loop will be broken
		for val := range bufferedChannel {
			fmt.Println(val)
		}

		fmt.Println("channel closed")
		done <- true
	}(bufferedChannel, done)

	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	bufferedChannel <- 4
	bufferedChannel <- 5
	bufferedChannel <- 6
	bufferedChannel <- 7
	bufferedChannel <- 8
	bufferedChannel <- 9

	close(bufferedChannel)

	//The program will wait until done is true
	<-done

	fmt.Println("done chan is true and the program is terminating now")
}
