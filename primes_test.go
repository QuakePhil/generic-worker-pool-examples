package main

import (
	"fmt"
	"github.com/quakephil/generic-worker-pool"
)

func ExamplePrimesWorker() {
	result := primes().Wait(1)
	fmt.Println("Found", result, "primes")
	// Output: Found 664579 primes
}

func ExamplePrimesWorkers() {
	result := primes().Wait(1000)
	fmt.Println("Found", result, "primes")
	// Output: Found 664579 primes
}

type batch struct {
	start  int
	end    int
	primes int
}

func primes() pool.Pool[batch, int] {
	return pool.New[batch, int](
		func(in chan<- batch) {
			for i := 1; i <= 10000000; i += 100000 {
				in <- batch{i, i + 100000, 0}
			}
		},
		// count how many primes there are in a given batch
		func(i batch) batch {
			for j := i.start; j < i.end; j++ {
				if isPrime(j) {
					i.primes++
				}
			}
			return i
		},
		// compute grand total
		func(results <-chan batch) int {
			count := 0
			for batch := range results {
				count += batch.primes
			}
			return count
		},
	)
}

// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
