package commands

import (
	"shell-analyzer/m/data"
	"shell-analyzer/m/tui"
	"strings"
)

func CheckTopOptions() {

	cmd := data.LineFields[0]
	if len(data.LineFields) > 1 && cmd == "sudo" {
		cmd = data.LineFields[1]
	}
	if len(data.LineFields) > 2 && strings.Contains(data.LineFields[1], "-") {
		cmd = data.LineFields[1]
	}

	tui.UnsortedTopCmds[cmd]++

}

func CheckSystemCommands() {
	if len(data.LineFields) > 1 && data.LineFields[0] != "sudo" {

		switch option := data.LineFields[0]; option {

		case "flatpak":
			tui.UnsortedSystemCmds["Flatpak Overall"]++

		case "apt":
			tui.UnsortedSystemCmds["APT Overall"]++

		case "dnf":
			tui.UnsortedSystemCmds["DNF Overall"]++

		case "yay":
			tui.UnsortedSystemCmds["Yay Overall"]++

		case "brew":
			tui.UnsortedSystemCmds["Brew Overall"]++
		}
	}
}
