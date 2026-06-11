package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, errors.New("Error: could not locate user home directory")
	}
	configPath := homeDir + configFileName

	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("Error: could not read app config file at path '%s'\n", configPath)
	}
	defer configFile.Close()

	cfg := Config{}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&cfg)

	return cfg, nil
}
