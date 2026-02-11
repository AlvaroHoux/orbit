package transport

import (
	"fmt"
	"net"
)

func Listen(data []byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}