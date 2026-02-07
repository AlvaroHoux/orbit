package config

import (
	"log"
	"os"
	"path/filepath"
	"encoding/json"
)

type GlobalConfig struct {
	DeviceID string `json:"deviceId"`
	Vaults []Vault `json:"vaults"`
}

type Vault struct {
	Id string `json:"id"`
	Path string `json:"path"`
}

// LoadGlobal tenta carregar. Se não existir, cria e retorna o padrão
func LoadGlobal() (*GlobalConfig, error) {
	configPath, err := getHomePath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		return createDefaultConfig(configPath)
	}

	if err != nil {
		return nil, err
	}

	var cfg GlobalConfig
	if err := json.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// createDefaultConfig gera, salva e retorna a config padrão
func createDefaultConfig(configPath string) (*GlobalConfig, error) {
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-device"
	}

	defaultConfig := GlobalConfig{
		DeviceID: hostname,
		Vaults: []Vault{},
	}

	jsonData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		log.Fatalf("Não foi possível criar arquivo de configuração: %v", err)
	}

	if err := os.WriteFile(configPath, jsonData, 0644); err != nil {
		return nil, err
	}

	return &defaultConfig, nil
}

func getHomePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".orbit", "registry.json"), nil
}
