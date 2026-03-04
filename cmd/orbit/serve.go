package main

import "orbit/internal/transport"

func runServe(args []string) {
	go transport.StartServer()
	go transport.StartDiscoveryListner()
	select {}
}
