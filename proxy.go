package main

import (
	"fmt"
	"net"
)

func proxy(source string, target string) {
	fmt.Println("TunnelMe Proxy")
	ln, _ := net.Listen("tcp", source)
	// accept connection on port
	for {
		conn, _ := ln.Accept()
		go func() {
			clientConn, _ := net.Dial("tcp", target)
			pipe(conn, clientConn)
		}()
	}
}
