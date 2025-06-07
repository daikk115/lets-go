package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	// fmt.Println("Received BC:", <-ch1)

	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)

		// if enable this, it will print "Timeout!" after 1 second
		// case <-time.After(1 * time.Second):
		// 	fmt.Println("Timeout!")

		// if enable this, it will immediately print "no channel ready"
		// default:
		// 	fmt.Println("no channel ready")
	}
}
