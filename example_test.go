package main

import (
	"fmt"
	"github.com/quakephil/generic-worker-pool"
)

// 1 worker takes 1 second
func ExampleSleepWorker() {
	result := pool.New[int, int](SleepWorker{}, sleepTest).Wait(1)
	fmt.Println("Slept", result, "times")
	// Output: Slept 10 times
}

// 10 workers take 1/10th of a second
func ExampleSleepWorkers() {
	result := pool.New[int, int](SleepWorker{}, sleepTest).Wait(10)
	fmt.Println("Slept", result, "times")
	// Output: Slept 10 times
}

func ExamplePrimesWorker() {
	result := pool.New[PrimesRange, int](PrimesWorker{}, primesTest).Wait(1)
	fmt.Println("Found", result, "primes")
	// Output: Found 1270607 primes
}

func ExamplePrimesWorkers() {
	result := pool.New[PrimesRange, int](PrimesWorker{}, primesTest).Wait(1000)
	fmt.Println("Found", result, "primes")
	// Output: Found 1270607 primes
}

// Test inputs

// 10 units of sleep = 1 second.
func sleepTest(in chan int) {
	for n := 1; n <= 10; n++ {
		in <- 100
	}
}

func primesTest(in chan PrimesRange) {
	for i := 1; i <= 20000000; i += 100000 {
		in <- PrimesRange{i, i + 100000, 0}
	}
}
