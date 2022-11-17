// Sleeps n milliseconds at a time, i.e. pretending to do some expensive task across the network
package main

import "time"

type SleepWorker struct{}

func (w SleepWorker) Process(n int) int {
	time.Sleep(time.Duration(n) * time.Millisecond) // placeholder for busywork
	return n
}

func (w SleepWorker) Output(processed chan int) (count int) {
	for _ = range processed {
		count++
	}
	return
}
