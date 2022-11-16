package main

import "net/url"

func (w Worker) Process(i Work) Work {
	host, err := url.Parse(i.registrarWebsite)
	check(err)
	i.ping = pingHost(host.Hostname(), w.timeout)
	return i
}
