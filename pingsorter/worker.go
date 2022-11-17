package main

import "time"

// Worker
type Worker struct {
	ascending bool
	timeout   time.Duration
	result    chan time.Duration
}

func PingSorter() (w Worker) {
	w.ascending = false
	w.timeout = time.Second
	w.result = make(chan time.Duration, 1)
	return
}

// Work
type Work struct {
	registrar        string
	iana             string
	location         string
	contact          string
	registrarWebsite string
	ping             time.Duration
}

func WorkFromRecord(record []string) Work {
	return Work{
		record[0],
		record[1],
		record[2],
		record[3],
		record[4],
		0,
	}
}

// helpers
func check(e error) {
	if e != nil {
		panic(e)
	}
}
