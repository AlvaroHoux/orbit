package main

import (
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "io"
	"fmt"
)

func encryptString(fonte string, senha string) ([]byte, error) {
	if senha == "" {
		return nil, fmt.Errorf("senha vazia")
	}
	
	key := sha256.Sum256([]byte(senha))
	block, err := aes.NewCipher(key[:])
	
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil{return nil, err}

	plaintext := []byte(fonte)
	
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil

}

func decryptString(ciphertext []byte, senha string) (string, error) {
	if senha == ""{
		return "", fmt.Errorf("senha vazia")
	}

	key := sha256.Sum256([]byte(senha))
	block, err := aes.NewCipher(key[:])
	
	if err != nil {
		return "", err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < gcm.NonceSize(){
		return "", fmt.Errorf("tamanho incorreto do ciphertext")
	}

	extractedNonce := ciphertext[:nonceSize]

	ciphertextReal := ciphertext[nonceSize:]

	plaintextBytes, err := gcm.Open(nil, extractedNonce, ciphertextReal, nil)
	if err != nil {
		return "", err
	}
	
	return string(plaintextBytes), nil
}
