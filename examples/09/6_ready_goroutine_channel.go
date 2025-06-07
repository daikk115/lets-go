package main

import (
	"fmt"
	"time"
)

var c chan string

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	c <- fmt.Sprintf("%s is ready after %v!", w, sec)
}

func main() {
	c = make(chan string)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	fmt.Println(<-c)
	fmt.Println(<-c)
}
