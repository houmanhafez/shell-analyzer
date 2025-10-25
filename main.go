package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"shell-analyzer/m/commands"
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
		defer file.Close()

		historyScanner := bufio.NewScanner(file)
		for historyScanner.Scan() {
			rawLine := strings.TrimSpace(historyScanner.Text())
			if rawLine == "" {
				continue
			}

			lines := strings.Split(rawLine, "\n")
			for _, singleLine := range lines {
				if singleLine == "" {
					continue
				}

				if strings.HasPrefix(singleLine, ":") && strings.Contains(singleLine, ";") {
					parts := strings.SplitN(singleLine, ";", 2)
					if len(parts) < 2 {
						continue
					}

					timestamp := strings.TrimSpace(strings.SplitN(parts[0], ":", 3)[1])
					unixInt, err := strconv.ParseInt(timestamp, 10, 64)
					if err != nil {
						log.Println("Invalid timestamp:", err)
						continue
					}

					data.CommandTime = time.Unix(unixInt, 0)
					singleLine = parts[1]

					if time.Since(data.CommandTime) <= 24*time.Hour && time.Now().After(data.CommandTime) {
						tui.OtherCount["Commands today"]++
					}
				}

				data.LineFields = strings.Fields(singleLine)
				if len(data.LineFields) == 0 {
					continue
				}

				commands.CheckGitCommits()

				cmd := data.LineFields[0]
				if cmd == "sudo" && len(data.LineFields) > 1 {
					cmd = data.LineFields[1]
				}

				tui.CommandCount[cmd]++
			}
		}
	}

	tui.CreateTextView()
}
