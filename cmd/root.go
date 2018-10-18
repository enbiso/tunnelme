package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tunnelme",
	Short: "Tunnel Me is a TCP tunneling tool",
	Long: `A Fast and Flexible TCP tunning tool to expose 
				  localhost's TCP ports (behind NAT) to internet
				  Complete documentation is available at https://www.enbiso.com/tunnelme`,
	Run: func(cmd *cobra.Command, args []string) {
		client()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TunnelMe",
	Long:  `All software has versions. This is TunnelMe's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tunnel Me v0.0.1 - beta")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(proxyCmd)
	rootCmd.AddCommand(serverCmd)
}

//Execute Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
