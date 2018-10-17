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
		server()
	} else if args[0] == "client" {
		client()
	} else if args[0] == "proxy" {
		proxy(":80", ":90")
	} else {
		usage()
	}
}

func usage() {
	log("Usage: tunnelme server (options)")
	log("Usage: tunnelme client (options)")
	log("Usage: tunnelme proxy (options)")
	os.Exit(-1)
}
