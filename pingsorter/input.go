package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

var source = "www.icann.org/en/accredited-registrars/Accredited-Registrars-202211161048.csv"
var linesToSkip = 1

func getRecords() (records [][]string) {
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

func input(in chan Work) {
	records := getRecords()
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
