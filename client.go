package main

import (
	"bufio"
	"fmt"
	"net"
)

// Client start point
func client() {
	fmt.Println("TunnelMe Client")
	ctrlConn, _ := net.Dial("tcp", "127.0.0.1:9001")
	remoteAddress, _ := bufio.NewReader(ctrlConn).ReadString('\n')
	fmt.Println("C: Remote - " + remoteAddress)
	for {
		id, _ := bufio.NewReader(ctrlConn).ReadString('\n')
		fmt.Println("C: connection ID recieved - " + id)

		dataConn, _ := net.Dial("tcp", "127.0.0.1:9002")
		fmt.Fprintf(dataConn, id)
		serviceConn, _ := net.Dial("tcp", "127.0.0.1:3306")
		fmt.Println("C: Piping")
		go pipe(dataConn, serviceConn)
	}
}
