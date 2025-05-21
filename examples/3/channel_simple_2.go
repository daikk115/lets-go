package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		messages <- "delay ping 5 second"
	}()
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		messages <- "delay ping 1 second"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
