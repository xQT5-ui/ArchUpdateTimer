package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Command struct {
		Cmnd         string   `yaml:"exec_value"`
		CmdShellFlag bool     `yaml:"cmd_shell_flag"`
		CmdShellExec string   `yaml:"cmd_shell_exec"`
		Emulation    []string `yaml:"emulation_terminal"`
	} `yaml:"command"`
	Days struct {
		Range    []int `yaml:"range"`
		RangeFlg bool  `yaml:"range_flg"`
	} `yaml:"days"`
	Pswrd string `yaml:"password"`
	Hash  string `yaml:"hash"`
}

func LoadConfig() Config {
	// Create new struct
	var config Config

	// Get executable path
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("ошибка получения пути исполняемого файла: %v", err)
		return config
	}

	// Get executable dir
	execDir := filepath.Dir(execPath)

	// Set config path
	filename := filepath.Join(execDir, "config", "config.yaml")

	// Check if file exists
	_, err = os.Stat(filename)
	if err != nil {
		log.Fatalf("файл конфигурации '%s' не существует: %v", filename, err)
		return config
	}

	// Check if file is readable
	_, err = os.Open(filename)
	if err != nil {
		log.Fatalf("файл конфигурации '%s' недоступен для чтения: %v", filename, err)
		return config
	}

	// Read config file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ошибка чтения файла '%s': %v", filename, err)
		return config
	}

	// Parsing YAML to struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("ошибка парсинга файла '%s': %v", filename, err)
		return config
	}

	return config
}
