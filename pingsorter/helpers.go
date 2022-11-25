package main

import (
	"encoding/csv"
	"github.com/go-ping/ping"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func getUniqueRecords(in chan<- registrar) {
	records := getRecords()
	checkDupes := make(map[string]bool)

	for _, record := range records {
		if config.linesToSkip > 0 {
			config.linesToSkip--
			continue
		}
		domainName := strings.ToLower(record[4])
		if _, exists := checkDupes[domainName]; !exists {
			checkDupes[domainName] = true
			in <- registrarFromRecord(record)
		}
	}
}

func getRecords() (records [][]string) {
	log.Println("reading", config.source)
	f, err := os.Open(config.source)
	check(err)
	defer f.Close()

	input := csv.NewReader(f)
	input.LazyQuotes = true
	records, err = input.ReadAll()
	check(err)
	return
}

func registrarFromRecord(record []string) registrar {
	return registrar{
		record[0],
		record[1],
		record[2],
		record[3],
		record[4],
		0,
	}
}

func pingHost(host string) time.Duration {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Println(err)
	}
	pinger.Count = 1
	pinger.Timeout = config.timeout
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		log.Println(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	return stats.AvgRtt
}

func sorter(pings []registrar) {
	sort.SliceStable(pings, func(i, j int) bool {
		if config.ascending {
			return pings[i].ping < pings[j].ping
		}
		return pings[i].ping > pings[j].ping
	})
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
