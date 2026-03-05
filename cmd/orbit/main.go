package main

import (
	"fmt"
	"log"
	"orbit/internal/config"
	"os"
)

func main() {
	globalConfig, err := config.LoadGlobal()
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Uso: orbit [serve|join|add|list]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "serve":
		runServe(args)
	case "join":
		runJoin(args, *globalConfig)
	case "stop":
		runToggleVault(args, *globalConfig, false)
	case "start":
		runToggleVault(args, *globalConfig, true)	
	case "list":
		runList(args)
	case "add":
		runAdd(args)
	default:
		fmt.Println("Comando desconhecido.")
		os.Exit(1)
	}
}
