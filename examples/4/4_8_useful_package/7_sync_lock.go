/*
- Random delay in Incrase --> simulate long running process
- time.Sleep in the end of multiRoutineLock to wait for all goroutines to finish --> Adjust to be long enough
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// https://go.dev/tour/concurrency/9

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	lock  sync.Mutex
	value int
}

// Increase the counter for the given key.
func (c *SafeCounter) Increase(routine int) {
	delay := rand.Intn(3000) // random number between 0–2999 milliseconds
	time.Sleep(time.Duration(delay) * time.Millisecond)

	c.lock.Lock()
	// Lock so only one goroutine at a time can access the map c.value.
	fmt.Println("Lock for increase in routine", routine, "with value", c.value)
	c.value++
	c.lock.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.lock.Lock()
	// Lock so only one goroutine at a time can access the map c.value.
	defer c.lock.Unlock()
	return c.value
}

func multiRoutineLock() {
	fmt.Println("\n------------------ multiRoutineLock ------------------")
	c := SafeCounter{value: 0}
	for i := 0; i < 100; i++ {
		go c.Increase(i)
	}

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println(c.Value())
}

func main() {
	multiRoutineLock()
}
