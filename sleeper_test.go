// Sleeps 100 milliseconds at a time
package main

import (
	"fmt"
	"github.com/quakephil/generic-worker-pool"
	"time"
)

// 1 worker takes 1 second
func ExampleSleepWorker() {
	result := sleeper().Wait(1)
	fmt.Println("Slept", result, "times")
	// Output: Slept 10 times
}

// 10 workers take 1/10th of a second
func ExampleSleepWorkers() {
	result := sleeper().Wait(10)
	fmt.Println("Slept", result, "times")
	// Output: Slept 10 times
}

func sleeper() pool.Pool[int, int] {
	return pool.New[int, int](
		// 10 units of sleep = 1 second
		func(in chan<- int) {
			for n := 1; n <= 10; n++ {
				in <- 100 // sleep 100 ms at a time
			}
		},
		// pretend to do some expensive task across the network
		func(n int) int {
			time.Sleep(time.Duration(n) * time.Millisecond)
			return n
		},
		// collect results (optional)
		func(results <-chan int) (count int) {
			for _ = range results {
				count++
			}
			return // how many times we slept
		},
	)
}
