package commands

import (
	"shell-analyzer/m/data"
	"shell-analyzer/m/tui"
	"time"
)

func CheckGitCommands() {

	if len(data.LineFields) > 1 && data.LineFields[0] == "git" {

		switch option := data.LineFields[1]; option {

		case "commit":
			tui.UnsortedGitCmds["Commits Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Commits Today"]++
			}
		case "pull":
			tui.UnsortedGitCmds["Pulls Overall"]++
		}
	}
}
