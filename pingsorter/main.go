package main

import (
	"fmt"
	"github.com/quakephil/generic-worker-pool"
	"log"
	"net/url"
	"time"
)

var config struct {
	source      string
	linesToSkip int
	ascending   bool
	timeout     time.Duration
}

type registrar struct {
	name     string
	iana     string
	location string
	contact  string
	website  string
	ping     time.Duration
}

func init() {
	config.source = "www.icann.org/en/accredited-registrars/Accredited-Registrars-202211161048.csv"
	config.linesToSkip = 1
	config.ascending = false
	config.timeout = time.Second
}

func main() {
	start := time.Now()
	total := pingsorter().Wait(100)
	log.Println("total pings:", total, "total runtime:", time.Since(start))
}

func pingsorter() pool.Pool[registrar, time.Duration] {
	return pool.New[registrar, time.Duration](
		// input
		func(in chan<- registrar) {
			getUniqueRecords(in)
		},
		// worker
		func(i registrar) registrar {
			host, err := url.Parse(i.website)
			check(err)
			i.ping = pingHost(host.Hostname())
			return i
		},
		// output
		func(processed <-chan registrar) (total time.Duration) {
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
		},
	)
}
