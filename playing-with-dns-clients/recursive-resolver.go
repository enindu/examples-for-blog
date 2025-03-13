package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func RecursiveResolver() {
	var (
		domain            string
		recursiveResolver string
	)

	// Point 1
	flag.StringVar(&domain, "d", "", "Domain name to lookup")
	flag.StringVar(&recursiveResolver, "rr", "8.8.8.8", "IP address of recursive resolver to use")
	flag.Parse()

	if domain == "" {
		flag.PrintDefaults()
		return
	}

	// Point 2
	client := dns.Client{}

	// Point 3
	message := dns.Msg{}

	message.SetQuestion(dns.Fqdn(domain), dns.TypeNS)

	// Point 4
	response, _, err := client.Exchange(&message, net.JoinHostPort(recursiveResolver, "53"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response.String())
}
