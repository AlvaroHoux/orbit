package main

import (
	"fmt"
	"orbit/internal/config"
	"orbit/internal/transport"
	"os"
)

func runJoin(args []string, globalConfig config.GlobalConfig) {
	if len(os.Args) < 3 {
		fmt.Println("Informe o ID do vault!")
		return
	}

	vault := args[0]

	ip := transport.FindPeers()
	if ip != "" {
		transport.Connect(ip, globalConfig.DeviceID, vault)
	}
}
