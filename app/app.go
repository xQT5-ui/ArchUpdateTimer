package main

import (
	"time"

	"app.go/app/config"
	"app.go/app/lib"
)

func main() {
	// Load and read config
	config := config.LoadConfig()

	// Get decrypt password
	decryptPassword := lib.GetDecryptPassword(config.Pswrd, config.Hash)

	// Get today's day of week
	nowDayOfWeek := lib.GetDayOfWeek()

	// Execute command
	// 0 = Sunday
	if nowDayOfWeek == time.Weekday(config.Days[0]) || nowDayOfWeek == time.Weekday(config.Days[1]) || nowDayOfWeek == time.Weekday(config.Days[2]) {
		if config.Command.CmdShellFlag {
			lib.RunCmdInTerminal("echo " + decryptPassword + " | " + config.Command.CmdShellExec + " && " + config.Command.Cmnd)
		} else {
			lib.RunCmdInTerminal(config.Command.Cmnd)
		}
	}
}
