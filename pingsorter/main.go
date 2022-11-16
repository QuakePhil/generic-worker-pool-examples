package main

import (
	"github.com/quakephil/generic-worker-pool"
	"log"
	"time"
)

// go run . 1>pings.txt 2>errors.log
func main() {
	var total time.Duration
	start := time.Now()

	worker := PingSorter()

	go func() {
		for result := range worker.result {
			total += result
		}
	}()

	pool.New[Work](worker).Wait(100)

	log.Println("total pings:", total, "total runtime:", time.Since(start))
}
