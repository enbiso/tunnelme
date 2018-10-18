package cmd

import (
	"bufio"
	"fmt"
	"net"

	"github.com/enbiso/tunnelme/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Execute tunneling client",
	Long:  `Excute a tunnel client inside the NAT where the service can be accessed`,
	Run: func(cmd *cobra.Command, args []string) {
		client()
	},
}

// Client start point
func client() {
	log.Info("TunnelMe Client")
	ctrlConn, _ := net.Dial("tcp", "127.0.0.1:9001")
	remoteAddress, _ := bufio.NewReader(ctrlConn).ReadString('\n')
	log.Info("Remote Address: " + remoteAddress)
	for {
		id, _ := bufio.NewReader(ctrlConn).ReadString('\n')
		log.Debug("Connection ID " + id)

		dataConn, _ := net.Dial("tcp", "127.0.0.1:9002")
		fmt.Fprintf(dataConn, id)
		serviceConn, _ := net.Dial("tcp", "127.0.0.1:3306")
		go utils.Pipe(dataConn, serviceConn)
		log.Debug("Connected to source service")
	}
}
