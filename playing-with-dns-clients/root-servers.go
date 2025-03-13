package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func RootServers() {
	var domain string

	flag.StringVar(&domain, "d", "", "Domain name to lookup")
	flag.Parse()

	if domain == "" {
		flag.PrintDefaults()
		return
	}

	client := dns.Client{}
	message := dns.Msg{}

	message.SetQuestion(dns.Fqdn(domain), dns.TypeNS)

	// Point 1
	rootServers := []string{
		"198.41.0.4",
		"170.247.170.2",
		"192.33.4.12",
		"199.7.91.13",
		"192.203.230.10",
		"192.5.5.241",
		"192.112.36.4",
		"198.97.190.53",
		"192.36.148.17",
		"192.58.128.30",
		"193.0.14.129",
		"199.7.83.42",
		"202.12.27.33",
	}

	// Point 2
	for _, rootServer := range rootServers {
		response, _, err := client.Exchange(&message, net.JoinHostPort(rootServer, "53"))
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", response.String())
	}
}
