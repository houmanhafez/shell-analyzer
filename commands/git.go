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
			tui.UnsortedGitCmds["Git Commits Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Commits Today"]++
			}

		case "pull":
			tui.UnsortedGitCmds["Git Pulls Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Pulls Today"]++
			}

		case "push":
			tui.UnsortedGitCmds["Git Pushes Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Pushes Today"]++
			}

		case "add":
			tui.UnsortedGitCmds["Git Adds Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Adds Today"]++
			}

		case "status":
			tui.UnsortedGitCmds["Git Statuses Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Statuses Today"]++
			}

		case "checkout":
			if data.LineFields[2] == "-b" {
				tui.UnsortedGitCmds["Git Branches Created Overall"]++
			}
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmds["Git Branch Switches Today"]++
			}
		}
	}
}
