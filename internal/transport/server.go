package transport

import (
	"encoding/json"
	"fmt"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", PORT))
	if err != nil {
		fmt.Printf("Error ao iniciar Server TCP: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP ouvindo na porta", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error no Accept: %v", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	remoteIp := conn.RemoteAddr().String()
	fmt.Printf("Cliente conectado: %s\n", remoteIp)

	decoder := json.NewDecoder(conn)

	for {
		var msg Packet

		err := decoder.Decode(&msg)
		if err != nil {
			return
		}

		switch msg.Type {
		case MsgHandshake:
			var data HandshakePayload
			json.Unmarshal([]byte(msg.Payload), &data)

			fmt.Printf("Recebi um handshake! Usuário %s querendo vault %s\n", data.DeviceId,  data.VaultID)
		}
	}
}

func StartDiscoveryListner() {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%s", PORT))
	conn, _ := net.ListenUDP("udp", addr)
	if conn != nil {
		defer conn.Close()
	}

	fmt.Printf("Discovery UDP ouvindo na porta %s...\n", PORT)
	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		message := string(buffer[:n])

		if message == DISCOVERY_MAGIC_REQ {
			fmt.Printf("Peer encontrado via UDP: %s\n", remoteAddr.IP.String())
			_, err := conn.WriteToUDP([]byte(DISCOVERY_MAGIC_RES), remoteAddr)
			if err != nil {
				fmt.Printf("Erro ao responder UDP: %v\n", err)
			}
		}
	}
}
