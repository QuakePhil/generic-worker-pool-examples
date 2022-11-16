package main

import (
	"fmt"
	"log"
	"sort"
)

func sorter(pings []Work, ascending bool) {
	sort.SliceStable(pings, func(i, j int) bool {
		if ascending {
			return pings[i].ping < pings[j].ping
		}
		return pings[i].ping > pings[j].ping
	})
}

func (w Worker) Output(out chan Work) {
	pings := []Work{}
	for o := range out {
		pings = append(pings, o)
	}

	sorter(pings, w.ascending)
	for _, o := range pings {
		if o.ping == 0 {
			log.Printf("ping %s: timeout\n", o.registrarWebsite)
		} else {
			w.result <- o.ping
			fmt.Printf("ping %s (%s): %v\n", o.registrarWebsite, o.location, o.ping)
		}
	}
	close(w.result)
}
