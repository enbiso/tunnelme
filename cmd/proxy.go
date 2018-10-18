package cmd

import (
	"net"

	"github.com/spf13/viper"

	"github.com/enbiso/tunnelme/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Execute tunneling proxy",
	Long:  "Excute a simple proxy to tunnel the service via a different port",
	Run: func(cmd *cobra.Command, args []string) {
		proxy(cmd.PersistentFlags().Lookup("source").Value.String(), cmd.PersistentFlags().Lookup("dest").Value.String())
	},
}

func proxyInit() {
	proxyCmd.PersistentFlags().StringP("source", "s", "127.0.0.1:80", "Source Address")
	proxyCmd.PersistentFlags().StringP("dest", "d", "0.0.0.0:90", "Destination Address")

	viper.BindPFlag("source", proxyCmd.PersistentFlags().Lookup("source"))
	viper.BindPFlag("dest", proxyCmd.PersistentFlags().Lookup("dest"))

	viper.SetDefault("source", "127.0.0.1:80")
	viper.SetDefault("dest", "0.0.0.0:90")
}

func proxy(source string, target string) {
	log.Info("TunnelMe Proxy")
	ln, err := net.Listen("tcp", target)
	if err != nil {
		log.Fatal(err)
	}
	// accept connection on port
	for {
		conn, _ := ln.Accept()
		go func() {
			clientConn, err := net.Dial("tcp", source)
			if err != nil {
				log.Error(err)
				conn.Close()
				return
			}
			utils.Pipe(conn, clientConn)
		}()
	}
}
