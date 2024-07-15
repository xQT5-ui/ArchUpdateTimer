package lib

import (
	"log"
	"os/exec"
)

func RunCmdInTerminal(emulation []string, cmd string) {
	// Execute "terminal" with command
	terminalCmd := exec.Command(emulation[0], emulation[1], cmd)

	// Get stdout
	stdoutStderr, err := terminalCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("ошибка создания stdout/error pipe: %v", err)
	}
	log.Printf("Прочитано: %s\n", stdoutStderr)
}
