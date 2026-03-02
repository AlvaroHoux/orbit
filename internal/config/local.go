package config

import (
    "fmt"
    "os"
    "path/filepath"
    "orbit/internal/crypto"
)

type FileInfo struct {
    Path string `json:"path"`
    Hash string `json:"hash"`
}

type Index struct {
    VaultName string     `json:"vault_name"`
    Files     []FileInfo `json:"files"`
}


func List() ([]string, error) { 
    arquivos, err := os.ReadDir(".")
    if err != nil {
        return nil, err
    }

    var vaults []string
    for _, a := range arquivos {
       if a.IsDir() && a.Name() != ".orbit" {
            vaults = append(vaults, a.Name())
       }
    }
    return vaults, nil 
}


func Add(vaults []string) error {
    // Garante que a pasta de destino existe
    if err := os.MkdirAll(".orbit", 0755); err != nil {
        return err
    }

    for _, nomeDoVault := range vaults {
        h, err := crypto.HashString(nomeDoVault)
        if err != nil {
            fmt.Printf("Erro ao gerar hash para %s: %v\n", nomeDoVault, err)
            continue // Pula para o próximo vault se este der erro
        }

        //Define o caminho: .orbit/hash... 
        caminho := filepath.Join(".orbit", h)

        //Escreve o arquivo. O conteúdo pode ser o próprio hash ou o nome original
        err = os.WriteFile(caminho, []byte(nomeDoVault), 0644)
        if err != nil {
            return err
        }
        fmt.Printf("Vault '%s' adicionado com hash %s\n", nomeDoVault, h)
    }
    return nil
}