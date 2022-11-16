// Counts how many primes there are in batches of "step" size
package main

import "fmt"

// PrimesState gets sent from Input() to Process() to Output()
type PrimesState struct {
	Start  int
	End    int
	Primes int
}

// worker has shared state e.g. handles, connections, settings, etc
type worker struct {
	max  int
	step int
	In   chan PrimesState
}

// New worker state (could also use new() or a struct literal)
func PrimesNew(max, step int) (w worker) {
	w.max = max
	w.step = step
	return
}

func PrimesNewWithChannel() (w worker) {
	w.In = make(chan PrimesState)
	return
}

// Input() generates PrimesState, e.g. reading from SQL, etc.
func (w worker) Input(in chan PrimesState) {
	if w.In != nil {
		// better way to chain channels? maybe in <- <-w.In
		for {
			if i, more := <-w.In; !more {
				return
			} else {
				in <- i
			}
		}
	} else {
		for i := 1; i <= w.max; i += w.step {
			in <- PrimesState{i, i + w.step, 0}
		}
	}
}

// Process() works on PrimesState, counting how many primes there are in a given range
func (w worker) Process(i PrimesState) PrimesState {
	for j := i.Start; j < i.End; j++ {
		if isPrime(j) {
			i.Primes += 1
		}
	}
	return i
}

// Output() works on PrimesState results from Process(), computing a total sum of primes
func (w worker) Output(out chan PrimesState) {
	count := 0
	for o := range out {
		count += o.Primes
	}
	fmt.Printf("Found %d primes\n", count)
}
