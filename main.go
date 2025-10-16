package main

import (
	"bufio"
	"fmt"
	"os"
	"shell-analyzer/m/data"
	"shell-analyzer/m/tui"
	"strings"
)

func main() {

	for _, filePath := range data.Files {
		file, err := os.Open(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Skipping %s, file does not exist.\n", filePath)
				continue
			}

			fmt.Println("Error opening file:", err)
			continue
		}

		historyScanner := bufio.NewScanner(file)
		for historyScanner.Scan() {
			line := strings.TrimSpace(historyScanner.Text())
			if line == "" {
				continue
			}

			lines := strings.SplitSeq(string(line), "\n")

			for line := range lines {
				fullLine := strings.Fields(line)

				if len(fullLine) == 0 {
					continue
				}
				if strings.HasPrefix(line, ":") && strings.Contains(line, ";") {
					parts := strings.SplitN(line, ";", 2)
					line = parts[1]
				}

				fields := strings.Fields(line)

				if len(fields) == 0 {
					continue
				}

				var option string
				if len(fields) > 1 {
					option = fields[1]
					if option == "commit" && len(fields) > 2 {
						tui.CommitCount["Commits so far"]++
					}
				} else {
					continue
				}

				cmd := fields[0]

				if cmd == "sudo" && len(fields) > 1 {
					cmd = fields[1]
				}

				tui.CommandCount[cmd]++
			}
		}
		defer file.Close()
	}

	tui.NewTextView()

}
