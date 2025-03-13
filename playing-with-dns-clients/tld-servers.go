package main

import (
	"flag"
	"fmt"
	"net"
	"slices"
	"strings"

	"github.com/miekg/dns"
)

func TLDServers() {
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
	rootServerResponse, _, err := client.Exchange(&message, net.JoinHostPort("198.41.0.4", "53"))
	if err != nil {
		panic(err)
	}

	// Point 2
	tldServerHostname := strings.TrimRight(strings.ReplaceAll(rootServerResponse.Ns[0].String(), rootServerResponse.Ns[0].Header().String(), ""), ".")

	// Point 3
	tldServerIPAddresses := []string{}

	ipAddresses, err := net.LookupIP(tldServerHostname)
	if err != nil {
		panic(err)
	}

	for _, ipAddress := range ipAddresses {
		tldServerIPAddress := ipAddress.To4().String()

		if tldServerIPAddress != "<nil>" && !slices.Contains(tldServerIPAddresses, tldServerIPAddress) {
			tldServerIPAddresses = append(tldServerIPAddresses, tldServerIPAddress)
		}
	}

	// Point 4
	tldServerResponse, _, err := client.Exchange(&message, net.JoinHostPort(tldServerIPAddresses[0], "53"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", tldServerResponse.String())
}
