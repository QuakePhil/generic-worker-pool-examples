package main

import (
	"github.com/quakephil/generic-worker-pool"
	"log"
	"time"
)

// ./run.sh
func main() {
	start := time.Now()
	total := pool.New[registrar, time.Duration](input, worker, output).Wait(100)
	log.Println("total pings:", total, "total runtime:", time.Since(start))
}
