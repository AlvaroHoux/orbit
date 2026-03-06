package crypto

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err
}

func DeriveKey(password string, salt []byte) []byte {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)

	key := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	return key
}