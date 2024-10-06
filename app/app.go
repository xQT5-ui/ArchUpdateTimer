package main

import (
	"time"

	"app.go/app/config"
	"app.go/app/lib"
	"app.go/app/lib/logger"
)

func main() {
	// Create logger
	logger := logger.NewLogger(nil)

	// Load and read config
	config := config.LoadConfig(logger)

	// Get decrypt password
	decryptPassword := lib.GetDecryptPassword(config.Pswrd, config.Hash, logger)

	// Get today's day of week
	nowDayOfWeek := lib.GetDayOfWeek()

	// Execute command
	if config.Days.RangeFlg {
		if nowDayOfWeek == time.Weekday(config.Days.Range[0]) || nowDayOfWeek == time.Weekday(config.Days.Range[1]) || nowDayOfWeek == time.Weekday(config.Days.Range[2]) {
			runCmdInTerminal(&config, logger, decryptPassword)
		}
	} else {
		runCmdInTerminal(&config, logger, decryptPassword)
	}
}
func runCmdInTerminal(config *config.Config, logger *logger.Logger, decryptPassword string) {
	var cmd string

	// Create command
	if config.Command.CmdShellFlag {
		cmd = "echo " + decryptPassword + " | " + config.Command.CmdShellExec + " && " + config.Command.Cmnd
	} else {
		cmd = config.Command.Cmnd
	}

	// Execute command
	lib.RunCmdInTerminal(config.Command.Emulation, cmd, config, logger)
}
