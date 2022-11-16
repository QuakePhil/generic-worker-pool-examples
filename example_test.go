package main

import "github.com/quakephil/generic-worker-pool"

// 10 units of sleep = 1 second.
// 1 worker takes 1 second
func ExampleSleepWorker() {
	pool.New[SleepState](SleepWorker{10}).Wait(1)
	// Output: done
}

// 10 workers take 1/10th of a second
func ExampleSleepWorkers() {
	pool.New[SleepState](SleepWorker{10}).Wait(10)
	// Output: done
}

func ExamplePrimesNew() {
	pool.New[PrimesState](PrimesNew(20000000, 100000)).Wait(1)
	// Output: Found 1270607 primes
}

func ExamplePrimesNewConcurrent() {
	pool.New[PrimesState](PrimesNew(20000000, 100000)).Wait(1000)
	// Output: Found 1270607 primes
}

func ExamplePrimesNewCustomInputChannel() {
	worker := PrimesNewWithChannel()
	go func() {
		for i := 1; i <= 20000000; i += 100000 {
			worker.In <- PrimesState{i, i + 100000, 0}
		}
		close(worker.In)
	}()
	pool.New[PrimesState](worker).Wait(1000)
	// Output: Found 1270607 primes
}
