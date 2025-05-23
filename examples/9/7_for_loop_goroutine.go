package main

import "fmt"

func shower(c chan int, quit chan bool) {

	// While loop for ever until quit is received
	for {
		select {
		case j := <-c:
			fmt.Printf("Got %d\n", j)
		case <-quit:
			fmt.Println("Quit signal received")
			break
		}
	}
}

func main() {
	// Init 2 channels: one for sending data, one for quitting
	ch := make(chan int)
	quit := make(chan bool)

	// Start the goroutine with loop checking for quit
	go shower(ch, quit)

	// Send 10 numbers to the channel
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- false // or true, does not matter
}
