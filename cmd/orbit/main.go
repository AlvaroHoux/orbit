package main

import (
	"flag"
	"fmt"
	"orbit/internal/transport"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "serve":
		go transport.StartServer()
		go transport.StartDiscoveryListner()
		select {}

	case "join":
		ip := transport.FindPeers()
		if ip != "" {
			transport.Connect(ip, []byte("Hello World"))
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Mostrando todos os cofres em orbita")

	case "add":
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args()
		fmt.Println(args[0])
	}
}
