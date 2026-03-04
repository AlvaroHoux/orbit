package transport

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func Connect(ip string, deviceId string, vault string) {
	target := net.JoinHostPort(ip, PORT)
	fmt.Printf("Tentando conectar TCP em %s...\n", target)

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Erro ao conectar no Peer: %v", err)
		return
	}
	defer conn.Close()

	handshake := HandshakePayload{
		DeviceId: deviceId,
		VaultID: vault,
	}
	payloadBytes, _ := json.Marshal(handshake)

	packet := Packet{
		Type: MsgHandshake,
		Payload: string(payloadBytes),
	}

	json.NewEncoder(conn).Encode(packet)

	fmt.Println("Dados enviados com sucesso via TCP!")
}

func FindPeers() string {
	remoteAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%s", PORT))
	localAddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:0")

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("Erro no DIalUDP: %v\n", err)
	}
	defer conn.Close()

	for i := 0; i < UDP_RETRIES; i++ {
		conn.WriteToUDP([]byte(DISCOVERY_MAGIC_REQ), remoteAddr)
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)

		if err == nil {
			msg := string(buffer[:n])
			if msg == DISCOVERY_MAGIC_RES {
				fmt.Printf("Peer encontrado no IP: %s\n", addr.IP.String())
				return addr.IP.String()
			}
		}
		time.Sleep(UDP_TIMEOUT)
	}

	fmt.Println("Nenhum peer encontrado.")
	return ""
}
