package transport

import (
	"fmt"
	"net"
	"time"
)

func Connect(ip string, data []byte) {
	target := net.JoinHostPort(ip, PORT)
	fmt.Printf("Tentando conectar TCP em %s...\n", target)

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Erro ao conectar no Peer: %v", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("Erro enviando dados: %v\n", err)
		return
	}
	fmt.Println("Dados enviados com sucesso via TCP!")
}

func FindPeers() (string) {
	remoteAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%s", PORT))
	localAddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:0")

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("Erro no DIalUDP: %v\n", err)
	}
	defer conn.Close()

	message := []byte("ORBIT")
	conn.Write(message)

	for i := 0; i < 3; i++ {
		conn.WriteToUDP([]byte("ORBIT"), remoteAddr)
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)

		if err == nil {
			msg := string(buffer[:n])
			if msg == "ORBIT HERE" {
				fmt.Printf("Peer encontrado no IP: %s\n", addr.IP.String())
				return addr.IP.String()
			}
		}
	}

	fmt.Println("Nenhum peer encontrado.")
	return ""
}