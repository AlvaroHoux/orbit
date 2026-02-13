package config

import (
	"encoding/json"
	"fmt"
	"hash"
	"os"
	"path/filepath"
)

type FileInfo struct {
    Path string `json:"path"`
    Hash string `json:"hash"`
}

type Index struct {
    VaultName string     `json:"vault_name"`
    Files     []FileInfo `json:"files"`
}


func add (hashes Index){
    
    var err error
    for i := 0; i < 5; i++ {
        err = os.MkdirAll(".orbit", 0755)
        if err == nil {
            fmt.Println("Pasta .orbit criada!")
            break 
        }
        fmt.Printf("Tentativa %d falhou, tentando novamente...\n", i+1)
    }   

    if err != nil {
        return fmt.Errorf("não foi possível criar a pasta .orbit após 5 tentativas: %w", err)
    }

    index := Index{
    VaultName: "MeuCofre",
    Files:     []FileInfo{}, // Começa vazio
    }  

    
    data, err := json.MarshalIndent(index, "", "    ")
    if err != nil {
        return fmt.Errorf("erro ao gerar JSON: %w", err)
    }
}