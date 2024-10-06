package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	lg "app.go/app/lib/logger"
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

func LoadConfig(lg *lg.Logger) Config {
	// Create new struct
	var config Config

	// Get executable path
	execPath, err := os.Executable()
	if err != nil {
		lg.Error(err, "Ошибка получения пути исполняемого файла:")
		return config
	}

	// Get executable dir
	execDir := filepath.Dir(execPath)

	// Set config path
	filename := filepath.Join(execDir, "config", "config.yaml")

	// Check if file exists
	_, err = os.Stat(filename)
	if err != nil {
		lg.Error(err, fmt.Sprintf("Файл конфигурации '%s' не существует:", filename))
		return config
	}

	// Check if file is readable
	_, err = os.Open(filename)
	if err != nil {
		lg.Error(err, fmt.Sprintf("Файл конфигурации '%s' недоступен для чтения:", filename))
		return config
	}

	// Read config file
	data, err := os.ReadFile(filename)
	if err != nil {
		lg.Error(err, fmt.Sprintf("Ошибка чтения файла '%s':", filename))
		return config
	}

	// Parsing YAML to struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		lg.Error(err, fmt.Sprintf("ошибка парсинга файла '%s':", filename))
		return config
	}

	lg.Info(fmt.Sprintf("Файл конфигурации '%s' успешно загружен", filename))

	return config
}
