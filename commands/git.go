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
			if IsToday(data.CmdTime) {
				tui.UnsortedGitCmdsDaily["Git Commits Today"]++
			}

		case "pull":
			tui.UnsortedGitCmdsTotal["Git Pulls Overall"]++
			if IsToday(data.CmdTime) {
				tui.UnsortedGitCmdsDaily["Git Pulls Today"]++
			}

		case "push":
			tui.UnsortedGitCmdsTotal["Git Pushes Overall"]++
			if IsToday(data.CmdTime) {
				tui.UnsortedGitCmdsDaily["Git Pushes Today"]++
			}

		case "add":
			tui.UnsortedGitCmdsTotal["Git Adds Overall"]++
			if IsToday(data.CmdTime) {
				tui.UnsortedGitCmdsDaily["Git Adds Today"]++
			}

		case "status":
			tui.UnsortedGitCmdsTotal["Git Statuses Overall"]++
			if IsToday(data.CmdTime) {
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

func IsToday(t time.Time) bool {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return t.After(today) && t.Before(today.Add(24*time.Hour))
}
