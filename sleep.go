// Sleeps n milliseconds at a time
package main

import "time"

func sleepWorker(n int) int {
	// pretend to do some expensive task across the network
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func sleepOutput(results <-chan int) (count int) {
	for _ = range results {
		count++
	}
	return // how many times we slept
}
