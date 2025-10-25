package commands

import (
	"shell-analyzer/m/data"
	"shell-analyzer/m/tui"
	"time"
)

func CheckGitCommits() {

	if len(data.LineFields) > 1 {
		option := data.LineFields[1]

		if option == "commit" {
			tui.CommitCount["Commits Overall"]++
			if time.Since(data.CommandTime) <= 24*time.Hour {
				tui.CommitCount["Commits Today"]++
			}
		}
	}
}
