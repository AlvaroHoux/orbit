package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"encoding/json"
)

type GlobalConfig struct {
	DeviceID string `json:"deviceId"`
	Vaults []Vault `json:"vaults"`
}

type Vault struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Active bool `json:"active"`
}

func SaveConfig(configChanges GlobalConfig) error {
	configPath, err := getHomePath()
	if err != nil {
		return err
	}

	configData, err := json.Marshal(configChanges)
	if err != nil {
		return nil
	}

	if err := os.WriteFile(configPath, []byte(configData), 0644); err != nil {
		return err
	}
	return nil
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

	var username string

	currentUser, err := user.Current()
	if err != nil {
		username = "unknown-user"
	} else {
		username = currentUser.Username
	}


	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-device"
	}

	defaultConfig := GlobalConfig{
		DeviceID: username + "@" + hostname,
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
	home, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "orbit", "registry.json"), nil
}
