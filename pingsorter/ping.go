package main

import (
	"github.com/go-ping/ping"
	"log"
	"time"
)

func pingHost(host string, timeout time.Duration) time.Duration {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Println(err)
	}
	pinger.Count = 1
	pinger.Timeout = timeout
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		log.Println(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	return stats.AvgRtt
}
