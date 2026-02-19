package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func hashString (alvo string) (string, error){
	if alvo == "" {
		return "", fmt.Errorf("alvo vazio")
	} 
	
	hash := sha256.New()
	io.WriteString(hash, alvo)
	return hex.EncodeToString(hash.Sum(nil)), nil
}