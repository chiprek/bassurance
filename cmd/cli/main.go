package main

import (
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/chiprek/bassurance/internal/cli_cmds"
	"github.com/pelletier/go-toml/v2"
)

var version = "v0.1.0"

func loadConfig() (*cli_cmds.Config, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("unable to find conf directory: %w", err)
	}
	configPath := filepath.Join(confDir, "/bassurance/.bassurance.toml")

	fileBytes, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &cli_cmds.Config{
				APIUrl: "http://localhost:8080/api/v1",
				Output: "json",
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var cfg cli_cmds.Config
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

	rootCmd := cli_cmds.NewRootCmd(cfg, version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
