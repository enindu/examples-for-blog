package main

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func Bonus() {
	servers := []string{
		"49.12.121.225",
		"49.12.121.101",
		"49.12.121.59",
		"49.12.121.38",
		"49.12.121.189",
		"49.12.121.216",
		"49.12.121.17",
		"49.12.121.32",
		"49.12.121.200",
		"49.12.121.144",
		"49.12.121.133",
		"49.12.121.57",
		"49.12.121.218",
		"49.12.121.147",
		"49.12.121.55",
	}

	client := dns.Client{}
	message := dns.Msg{}

	message.SetQuestion(dns.Fqdn("winterasiatours.com"), dns.TypeA)

	for _, server := range servers {
		fmt.Printf("Scanning %s", server)

		response, _, err := client.Exchange(&message, net.JoinHostPort(server, "53"))
		if err != nil {
			fmt.Printf("\r")
			fmt.Printf("%s -> %v\n", server, err)
			continue
		}

		fmt.Printf("\r")

		if len(response.Answer) > 0 {
			fmt.Printf("%s -> An open resolver\n", server)
		} else {
			fmt.Printf("%s -> Not an open resolver\n", server)
		}
	}
}
