package main

import (
	"fmt"
	"log"

	"app.go/app/config"
	"app.go/app/lib"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("ошибка загрузки конфигурации: %v\n", err)
	}

	nowDayOfWeek := lib.GetDayOfWeek()

	if nowDayOfWeek == 0 { // 0 = Sunday
		// echo pswrd | sudo -S sh -c 'echo Start updating system...\n' && yay -Syu --noconfirm && sudo flatpak repair && flatpak update -y && sudo flatpak remove --unused -y
		if config.Command.CmdShellFlag {
			lib.RunCmdInTerminal("echo " + fmt.Sprintf("%d", config.Pswrd) + " | " + config.Command.CmdShellExec + " && " + config.Command.Cmnd)
		} else {
			lib.RunCmdInTerminal(config.Command.Cmnd)
		}
	}
}
