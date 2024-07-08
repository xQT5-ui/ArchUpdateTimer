package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Command struct {
		Cmnd         string `yaml:"exec_value"`
		CmdShellFlag bool   `yaml:"cmd_shell_flag"`
		CmdShellExec string `yaml:"cmd_shell_exec"`
	} `yaml:"command"`
	Pswrd int `yaml:"password"`
}

func LoadConfig() (Config, error) {
	// Create new struct
	var config Config

	// Get executable path
	execPath, err := os.Executable()
	if err != nil {
		return config, fmt.Errorf("ошибка получения пути исполняемого файла: %w", err)
	}

	// Get config path
	execDir := filepath.Dir(execPath)

	// Set config path
	filename := filepath.Join(execDir, "config", "config.yaml")

	// Check if file exists
	_, err = os.Stat(filename)
	if err != nil {
		return config, fmt.Errorf("файл конфигурации '%s' не существует: %w", filename, err)
	}

	// Check if file is readable
	_, err = os.Open(filename)
	if err != nil {
		return config, fmt.Errorf("файл конфигурации '%s' недоступен для чтения: %w", filename, err)
	}

	// Read config file
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("ошибка чтения файла '%s': %w", filename, err)
	}

	// Parsing YAML to struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("ошибка парсинга файла '%s': %w", filename, err)
	}

	return config, nil
}
