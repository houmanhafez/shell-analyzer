package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

				var commandTime time.Time

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

					commandTime = time.Unix(unixInt, 0)
					singleLine = parts[1]

					if time.Since(commandTime) <= 24*time.Hour && time.Now().After(commandTime) {
						tui.OtherCount["Commands today"]++
					}
				}

				fields := strings.Fields(singleLine)
				if len(fields) == 0 {
					continue
				}

				cmd := fields[0]
				if cmd == "sudo" && len(fields) > 1 {
					cmd = fields[1]
				}

				if len(fields) > 1 {
					option := fields[1]

					if option == "commit" {
						tui.CommitCount["Commits Overall"]++
						if time.Since(commandTime) <= 24*time.Hour {
							tui.CommitCount["Commits Today"]++
						}
					}
				}

				tui.CommandCount[cmd]++
			}
		}
	}

	tui.CreateTextView()
}
