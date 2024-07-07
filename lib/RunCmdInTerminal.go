package lib

import (
	"log"
	"os/exec"
)

func RunCmdInTerminal(cmd string) {
	// Запуск внешней программы "terminal"
	terminalCmd := exec.Command("kgx", "-e", cmd)

	// Получаем stdout
	stdoutStderr, err := terminalCmd.CombinedOutput()
	if err != nil {
		log.Fatal("Error creating stdout/error pipe: ", err)
	}
	log.Printf("Readed: %s\n", stdoutStderr)
}
