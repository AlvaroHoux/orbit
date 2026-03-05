package main

import (
	"flag"
	"fmt"
	"orbit/internal/config"
	"os"
)

func runToggleVault(args []string, globalConfig config.GlobalConfig, active bool) {
	stopCmd := flag.NewFlagSet("stop", flag.ExitOnError)

	var id string
	var name string

	stopCmd.StringVar(&id, "id", "", "Id do vault")
	stopCmd.Parse(args)

	if args[0] != "-id" {
		name = args[0]
	}

	messages := [][]string{
		{"Iniciando vault...", "Vault iniciado com sucesso!", "O vault já está iniciado!"},
		{"Parando vault...", "Vault parado com sucesso!", "O vault já está parado!"},
	}

	var currentMessage []string

	if active {
		currentMessage = messages[0]
	} else {
		currentMessage = messages[1]
	}

	newGlobalConfig := globalConfig

	for i, vault := range newGlobalConfig.Vaults {
		if (name != "" && vault.Name == name) || (id != "" && vault.Id == id) {
			fmt.Println(currentMessage[0])
			fmt.Printf("ID: %s | Nome: %s | Caminho: %s\n", vault.Id, vault.Name, vault.Path)

			if vault.Active != active {
				newGlobalConfig.Vaults[i].Active = active
				fmt.Println(currentMessage[1])
			} else {
				fmt.Println(currentMessage[2])
			}
		}
	}

	if err := config.SaveConfig(newGlobalConfig); err != nil {
		fmt.Println("Ocorreu um erro ao salvar", err)
		os.Exit(1)
	}
}
