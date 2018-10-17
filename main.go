package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
	}
	if args[0] == "server" {
		Server()
	} else if args[0] == "client" {
		Client()
	} else {
		usage()
	}
}

func usage() {
	log("Usage: tunnelme server (options)")
	log("Usage: tunnelme client (options)")
	os.Exit(-1)
}
