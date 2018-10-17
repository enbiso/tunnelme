package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Server start point
func Server() {

	remoteStartPort := 8000

	dataConns := map[string]net.Conn{}

	fmt.Println("TunnelMe Server")
	ctrlLn, _ := net.Listen("tcp", ":9001")
	dataLn, _ := net.Listen("tcp", ":9002")

	go func() {
		for {
			dataConn, _ := dataLn.Accept()
			remortID, _ := bufio.NewReader(dataConn).ReadString('\n')
			fmt.Println("S: Data connection accepted")
			dataConns[strings.TrimSuffix(remortID, "\n")] = dataConn
		}
	}()

	for {
		ctrlConn, _ := ctrlLn.Accept()
		fmt.Println("S: Controller Accepted")

		remoteAddress := fmt.Sprintf(":%d", remoteStartPort)
		remoteStartPort++
		clientLn, _ := net.Listen("tcp", remoteAddress)
		fmt.Fprintf(ctrlConn, remoteAddress+"\n")

		go func() {
			for {
				clientConn, _ := clientLn.Accept()
				fmt.Println("S: Client Accepted")
				id := UUID()
				fmt.Fprintf(ctrlConn, id+"\n")
				fmt.Println("S: COMMAND Send")

				for dataConns[id] == nil {
					time.Sleep(200 * time.Millisecond)
				}

				dataConn := dataConns[id]

				fmt.Println("S: Piping")
				go Pipe(clientConn, dataConn)
			}
		}()
	}
}
