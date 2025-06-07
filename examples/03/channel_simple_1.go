package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		messages <- "delay ping 5 second" // send message to channel, this will block until someone reads from the channel (<-messages)
	}()
	go func() {
		time.Sleep(time.Duration(10) * time.Second)
		messages <- "delay ping 10 second" // send message to channel, this will block until someone reads from the channel (<-messages)
	}()

	fmt.Print("waiting for messages...\n")
	fmt.Println(<-messages) // Block until a message is sent to the channel, then read the message from the channel
	fmt.Print("Done 1 message...\n")
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
