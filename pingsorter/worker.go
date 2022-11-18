package main

import (
	"net/url"
	"time"
)

type registrar struct {
	name     string
	iana     string
	location string
	contact  string
	website  string
	ping     time.Duration
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

func worker(i registrar) registrar {
	host, err := url.Parse(i.website)
	check(err)
	i.ping = pingHost(host.Hostname())
	return i
}

// helpers
func check(e error) {
	if e != nil {
		panic(e)
	}
}
