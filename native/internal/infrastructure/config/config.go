package config

import (
	"encoding/json"
	"native/internal/infrastructure/logger"
	"os"
)

type Config struct {
	BaseDir string `json:"base_dir"`
}

func Load(path string) (*Config, error) {
	logger.Logger.Println("Loading config from:", path)
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Println("Error reading config file:", err)
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		logger.Logger.Println("Error unmarshalling config:", err)
		return nil, err
	}

	logger.Logger.Println("Config loaded:", cfg)
	return &cfg, nil
}
