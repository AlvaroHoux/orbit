package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"orbit/internal/transport"
)

func main()  {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "serve":
		transport.Server()
		
	case "join":
		fullString := strings.Join(os.Args[2:], " ")
		data := []byte(fullString)
		transport.Listen(data)

	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Mostrando todos os cofres em orbita")

	case "add":
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args()
		fmt.Println(args[0])
	}
}