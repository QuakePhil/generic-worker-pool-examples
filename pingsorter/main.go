package main

import (
	"github.com/quakephil/generic-worker-pool"
	"log"
	"time"
)

// ./run.sh
func main() {
	start := time.Now()
	total := pool.New[Work, time.Duration](input, PingSorter()).Wait(100)
	log.Println("total pings:", total, "total runtime:", time.Since(start))
}
