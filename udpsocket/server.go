package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	lAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	checkError(err)
	conn, err := net.ListenUDP("udp", lAddress)
	checkError(err)

	for {
		var (
			address string
			buf     [512]byte
		)
		_, rAddress, err := conn.ReadFrom(buf[0:])
		checkError(err)
		address = rAddress.String()
		fmt.Printf("%s: %s\n", address, string(buf[0:]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(1)
	}
}
