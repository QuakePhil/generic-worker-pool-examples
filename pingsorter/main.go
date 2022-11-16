package main

import "github.com/quakephil/generic-worker-pool"

// go run . 1>pings.txt 2>errors.log
func main() {
	worker := PingSorter()
	pool.New[Work](worker).Wait(100)
}
