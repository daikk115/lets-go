package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // This will read from the channel until it is closed
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Duration(1) * time.Second) // Simulate some work
		results <- j * 2
		// fmt.Println("Result:", <-results) --> invalid operation: cannot receive from send-only channel results (variable of type chan<- int)compilerInvalidReceive
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int, 5)    // Buffered channel
	results := make(chan int, 5) // Buffered channel

	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		time.Sleep(time.Duration(1) * time.Second) // Simulate jobs meanwhile 2 workers are working
		jobs <- j
		fmt.Printf("Job %d sent\n", j)
	}
	close(jobs) // Without this, the range loop in the worker goroutines will never exit until the end of main, result live below
	fmt.Println("All jobs sent to workers, waiting for results...")

	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
		fmt.Println("Current goroutines:", runtime.NumGoroutine())
		time.Sleep(time.Duration(1) * time.Second)
	}
}

// ubuntu@vm ~/G/l/e/3 (master)> go run ./channel_workers.go
// Job 1 sent
// Worker 2 processing job 1
// Worker 2 finished job 1
// Job 2 sent
// Worker 1 processing job 2
// Worker 1 finished job 2
// Job 3 sent
// Worker 2 processing job 3
// Worker 2 finished job 3
// Job 4 sent
// Worker 1 processing job 4
// Worker 1 finished job 4
// Job 5 sent
// All jobs sent to workers, waiting for results...
// Result: 2
// Current goroutines: 3
// Worker 2 processing job 5
// Worker 2 finished job 5
// Result: 4
// Current goroutines: 3
// Result: 6
// Current goroutines: 3
// Result: 8
// Current goroutines: 3
// Result: 10
// Current goroutines: 3

// ubuntu@vm ~/G/l/e/3 (master)> go run ./channel_workers.go
// Job 1 sent
// Worker 2 processing job 1
// Worker 2 finished job 1
// Job 2 sent
// Worker 1 processing job 2
// Worker 1 finished job 2
// Job 3 sent
// Worker 2 processing job 3
// Worker 2 finished job 3
// Job 4 sent
// Worker 1 processing job 4
// Worker 1 finished job 4
// Job 5 sent
// All jobs sent to workers, waiting for results...
// Result: 2
// Current goroutines: 3
// Worker 2 processing job 5
// Worker 2 finished job 5
// Result: 4
// Current goroutines: 1
// Result: 6
// Current goroutines: 1
// Result: 8
// Current goroutines: 1
// Result: 10
// Current goroutines: 1
