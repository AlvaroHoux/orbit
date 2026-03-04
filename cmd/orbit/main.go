package main

import (
	"flag"
	"fmt"
	"log"
	"orbit/internal/config"
	"orbit/internal/transport"
	"os"
)

func main() {
	globalConfig, err := config.LoadGlobal()
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %v\n", err)
		return
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	joinCmd := flag.NewFlagSet("join", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "serve":
		go transport.StartServer()
		go transport.StartDiscoveryListner()
		select {}

	case "join":
		if len(os.Args) < 3 {
			fmt.Println("Informe o ID do vault!")
			return;
		}

		joinCmd.Parse(os.Args[2:])
		vault := joinCmd.Arg(0)

		ip := transport.FindPeers()
		if ip != "" {
			transport.Connect(ip, globalConfig.DeviceID, vault)
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
