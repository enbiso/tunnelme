package cmd

import (
	"net"

	"github.com/enbiso/tunnelme/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Execute tunneling proxy",
	Long:  "Excute a simple proxy to tunnel the service via a different port",
	Run: func(cmd *cobra.Command, args []string) {
		proxy(":80", ":90")
	},
}

func proxy(source string, target string) {
	log.Info("TunnelMe Proxy")
	ln, _ := net.Listen("tcp", source)
	// accept connection on port
	for {
		conn, _ := ln.Accept()
		go func() {
			clientConn, _ := net.Dial("tcp", target)
			utils.Pipe(conn, clientConn)
		}()
	}
}
