// Sleeps 100 milliseconds at a time, i.e. pretending to do some expensive task across the network
package main

import (
	"fmt"
	"time"
)

type SleepState struct{}

type SleepWorker struct {
	Units int // how much input to generate
}

func (w SleepWorker) Input(in chan SleepState) {
	for n := 1; n <= w.Units; n++ {
		in <- SleepState{}
	}
}

func (w SleepWorker) Process(i SleepState) SleepState {
	time.Sleep(100 * time.Millisecond) // placeholder for busywork
	return i
}

func (w SleepWorker) Output(out chan SleepState) {
	for _ = range out {
	}
}

func (w SleepWorker) Done() {
	fmt.Println("done")
}
