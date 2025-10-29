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
			tui.UnsortedGitCmdsTotal["Git Commits Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Commits Today"]++
			}

		case "pull":
			tui.UnsortedGitCmdsTotal["Git Pulls Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Pulls Today"]++
			}

		case "push":
			tui.UnsortedGitCmdsTotal["Git Pushes Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Pushes Today"]++
			}

		case "add":
			tui.UnsortedGitCmdsTotal["Git Adds Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Adds Today"]++
			}

		case "status":
			tui.UnsortedGitCmdsTotal["Git Statuses Overall"]++
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Statuses Today"]++
			}

		case "checkout":
			if data.LineFields[2] == "-b" {
				tui.UnsortedGitCmdsTotal["Git Branches Created Overall"]++
			}
			if time.Since(data.CmdTime) <= 24*time.Hour {
				tui.UnsortedGitCmdsDaily["Git Branch Switches Today"]++
			}
		}
	}
}
