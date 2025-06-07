package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, sleep_time int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(time.Duration(sleep_time) * time.Second)
	c <- sum // send sum to c
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{-1, -2, 3}

	c := make(chan int)
	// Depend on sleep_time, the order of the output will be different
	go sum(s1, c, 1)
	go sum(s2, c, 5)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y)
}
