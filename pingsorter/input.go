package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

var linesToSkip = 1

func getRecords(source string) (records [][]string) {
	log.Println("reading", source)
	f, err := os.Open(source)
	check(err)
	defer f.Close()

	input := csv.NewReader(f)
	input.LazyQuotes = true
	records, err = input.ReadAll()
	check(err)
	return
}

func (w Worker) Input(in chan Work) {
	records := getRecords(w.source)
	checkDupes := make(map[string]bool)

	for _, record := range records {
		if linesToSkip > 0 {
			linesToSkip--
			continue
		}
		domainName := strings.ToLower(record[4])
		if _, exists := checkDupes[domainName]; !exists {
			checkDupes[domainName] = true
			in <- WorkFromRecord(record)
		}
	}
}
