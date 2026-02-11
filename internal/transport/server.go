package transport

import (
	"fmt"
	"net"
)

func Server() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor ouvindo na porta", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("> Error: %v\n", err)
			return
		}
		fmt.Printf("Recebido: %s\n", buffer[:n])
	}
}