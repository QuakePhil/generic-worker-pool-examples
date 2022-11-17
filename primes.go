// Counts how many primes there are in batches of "step" size
package main

// PrimesRange gets sent from Process() to Output()
type PrimesRange struct {
	Start  int
	End    int
	Primes int
}

// PrimesWorker has shared state e.g. handles, connections, settings, etc
type PrimesWorker struct{}

// Count how many primes there are in a given range
func (w PrimesWorker) Process(i PrimesRange) PrimesRange {
	for j := i.Start; j < i.End; j++ {
		if isPrime(j) {
			i.Primes += 1
		}
	}
	return i
}

// Compute grand total
func (w PrimesWorker) Output(processed chan PrimesRange) int {
	count := 0
	for o := range processed {
		count += o.Primes
	}
	return count
}
