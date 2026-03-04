package transport

import "time"

const (
	PORT        string        = "8080"
	BUFFER_SIZE int           = 32 * 1024

	DISCOVERY_MAGIC_REQ string = "ORBIT_V1_DISCOVERY_REQ"
	DISCOVERY_MAGIC_RES string = "ORBIT_V1_DISCOVERY_RES"

	UDP_TIMEOUT time.Duration = 2 * time.Second
	UDP_RETRIES int           = 3
)

type MessageType int

const (
	MsgHandshake MessageType = iota + 1
	MsgHandshakeAck
	MsgIndexRequest
	MsgIndexResponse
	MsgFileRequest
	MsgFileResponse
	MsgError
)

type Packet struct {
	Type    MessageType `json:"type"`
	Payload string      `json:"payload"`
}

type HandshakePayload struct {
	DeviceId string `json:"deviceId"`
	VaultID string `json:"vaultId"`
}