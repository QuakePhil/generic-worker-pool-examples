package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sorter(pings []registrar) {
	sort.SliceStable(pings, func(i, j int) bool {
		if config.ascending {
			return pings[i].ping < pings[j].ping
		}
		return pings[i].ping > pings[j].ping
	})
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
