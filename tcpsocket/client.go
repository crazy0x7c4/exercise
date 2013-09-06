package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, errConn := net.Dial("tcp", "127.0.0.1:8888")
	checkError(errConn)
	_, errWrite := conn.Write([]byte("hello world."))
	checkError(errWrite)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(1)
	}
}
