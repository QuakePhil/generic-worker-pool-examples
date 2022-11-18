// Counts primes
package main

type batch struct {
	start  int
	end    int
	primes int
}

// Count how many primes there are in a given batch
func primesWorker(i batch) batch {
	for j := i.start; j < i.end; j++ {
		if isPrime(j) {
			i.primes += 1
		}
	}
	return i
}

// Compute grand total
func primesOutput(results <-chan batch) int {
	count := 0
	for batch := range results {
		count += batch.primes
	}
	return count
}
