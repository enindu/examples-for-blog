package main

import (
	"bufio"
	"fmt"
	"net"
)

func TCPServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		// Point 1
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		// Point 2
		go func(c net.Conn) {
			// Point 3
		}(connection)
	}
}

func TCPServerFinal() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			reader := bufio.NewReader(c)

			// Point 1
			startLine, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s", startLine)

			// Point 2
			for {
				header, err := reader.ReadString('\n')
				if err != nil || header == "\r\n" {
					break
				}

				fmt.Printf("%s", header)
			}

			// Point 3
			writer := bufio.NewWriter(c)

			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString("Content-Length: 12\r\n")
			writer.WriteString("Content-Type: text/html; charset=utf-8\r\n")
			writer.WriteString("\r\n")
			writer.WriteString("Hello world!\n")
			writer.Flush()
		}(connection)
	}
}
