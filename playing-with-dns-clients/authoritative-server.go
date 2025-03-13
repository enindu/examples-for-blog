package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func AuthoritativeServer() {
	var (
		domain                  string
		authoritativeNameServer string
	)

	flag.StringVar(&domain, "d", "", "Domain name to lookup")
	flag.StringVar(&authoritativeNameServer, "ans", "", "IP address of authoritative name server to use")
	flag.Parse()

	if domain == "" || authoritativeNameServer == "" {
		flag.PrintDefaults()
		return
	}

	client := dns.Client{}
	message := dns.Msg{}

	message.SetQuestion(dns.Fqdn(domain), dns.TypeA)

	response, _, err := client.Exchange(&message, net.JoinHostPort(authoritativeNameServer, "53"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response.String())
}
