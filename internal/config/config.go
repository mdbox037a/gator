package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("Error: could not read app config file at path '%s'\n", configPath)
	}
	defer configFile.Close()

	cfg := Config{}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("Error: failed to decode config file at path %s\n", configPath)
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Error: could not locate user home directory")
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	newConfig, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return errors.New("Error: failed to parse current config in setting new user")
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, newConfig, 0600)
	if err != nil {
		return fmt.Errorf("Error: failed to write config to file %s\nDebug: %w\n", configPath, err)
	}

	return nil
}
