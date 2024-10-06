package lib

import (
	"fmt"
	"os/exec"

	conf "app.go/app/config"
	log "app.go/app/lib/logger"
)

func RunCmdInTerminal(emulation []string, cmd string, config *conf.Config, lg *log.Logger) {
	// Execute "terminal" with command
	terminalCmd := exec.Command(emulation[0], emulation[1], cmd)

	// Get stdout
	stdoutStderr, err := terminalCmd.CombinedOutput()
	if err != nil && !config.Command.CmdShellFlag {
		lg.Error(err, "Ошибка создания stdout/error pipe:")
		return
	}

	lg.Info(fmt.Sprintf("Прочитано: %s", string(stdoutStderr)))
}
