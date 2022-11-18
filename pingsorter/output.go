package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

var ascending bool = false

func sorter(pings []registrar) {
	sort.SliceStable(pings, func(i, j int) bool {
		if ascending {
			return pings[i].ping < pings[j].ping
		}
		return pings[i].ping > pings[j].ping
	})
}

func output(processed <-chan registrar) (total time.Duration) {
	pings := []registrar{} // TODO: use a better data structure to sort with
	for o := range processed {
		pings = append(pings, o)
	}

	sorter(pings)
	for _, o := range pings {
		if o.ping == 0 {
			log.Printf("ping %s: timeout\n", o.website)
		} else {
			total += o.ping
			fmt.Printf("ping %s (%s): %v\n", o.website, o.location, o.ping)
		}
	}
	return total
}
