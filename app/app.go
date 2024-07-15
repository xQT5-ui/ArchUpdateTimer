package main

import (
	"app.go/lib"
)

func main() {
	nowDayOfWeek := lib.GetDayOfWeek()

	if nowDayOfWeek == 0 { // 0 = Sunday
		lib.RunCmdInTerminal("echo 89230 | sudo -S sh -c 'echo Start updating system...\n' && yay -Syu --noconfirm && sudo flatpak repair && flatpak update -y && sudo flatpak remove --unused -y")
	}
}
