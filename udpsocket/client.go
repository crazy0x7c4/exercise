package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	checkError(err)
	conn.Write([]byte("hello world."))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(1)
	}
}
