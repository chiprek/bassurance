package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"

	_ "embed"
)

type Config struct {
	APIUrl string `toml:"api_url"`
	Output string `toml:"output"`
}

var version string

func loadConfig() (*Config, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("unable to find conf directory: %w", err)
	}
	configPath := filepath.Join(confDir, "/bassurance/.bassurance.toml")

	fileBytes, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				APIUrl: "http://localhost:8080/api/v1",
				Output: "json",
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var cfg Config
	err = toml.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TOML syntax: %w", err)
	}
	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("startup Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully loaded config\nTarget API: %s\n", cfg.APIUrl)

}
