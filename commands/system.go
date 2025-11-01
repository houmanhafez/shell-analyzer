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
			tui.UnsortedSystemCmdsTotal["Flatpak Overall"]++

		case "apt":
			tui.UnsortedSystemCmdsTotal["APT Overall"]++

		case "dnf":
			tui.UnsortedSystemCmdsTotal["DNF Overall"]++

		case "yay":
			tui.UnsortedSystemCmdsTotal["Yay Overall"]++

		case "brew":
			tui.UnsortedSystemCmdsTotal["Brew Overall"]++
		}
	}
}
