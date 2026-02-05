package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Mostrando todos os cofres em orbita")

	case "add":
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args()
		fmt.Println(args[0])
	}
}