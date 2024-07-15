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
	if config.Command.EmulationFlg {
		// 0 = Sunday
		if nowDayOfWeek == time.Weekday(config.Days.Range[0]) || nowDayOfWeek == time.Weekday(config.Days.Range[1]) || nowDayOfWeek == time.Weekday(config.Days.Range[2]) {
			if config.Command.CmdShellFlag {
				lib.RunCmdInTerminal(config.Command.Emulation, "echo "+decryptPassword+" | "+config.Command.CmdShellExec+" && "+config.Command.Cmnd)
			} else {
				lib.RunCmdInTerminal(config.Command.Emulation, config.Command.Cmnd)
			}
		}
	} else {
		if config.Command.CmdShellFlag {
			lib.RunCmdInTerminal(config.Command.Emulation, "echo "+decryptPassword+" | "+config.Command.CmdShellExec+" && "+config.Command.Cmnd)
		} else {
			lib.RunCmdInTerminal(config.Command.Emulation, config.Command.Cmnd)
		}
	}

}
