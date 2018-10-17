package main

import (
	"fmt"
	"net"
)

// Proxy port
func Proxy(source string, target string) {
	fmt.Println("TunnelMe Proxy")
	ln, _ := net.Listen("tcp", source)
	// accept connection on port
	for {
		conn, _ := ln.Accept()
		go func() {
			clientConn, _ := net.Dial("tcp", target)
			Pipe(conn, clientConn)
		}()
	}
}
