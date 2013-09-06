package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, errListen := net.Listen("tcp", "127.0.0.1:8888")
	checkError(errListen)
	for {
		conn, errConn := ln.Accept()
		checkError(errConn)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		var buf [512]byte
		_, errRead := conn.Read(buf[0:])
		if errRead != nil {
			if errRead == io.EOF {
				fmt.Println("End Of File")
				break
			} else {
				os.Exit(1)
			}
		}
		fmt.Println(string(buf[0:]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(1)
	}
}
