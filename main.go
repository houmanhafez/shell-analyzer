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

		historyScanner := bufio.NewScanner(file)
		for historyScanner.Scan() {
			line := strings.TrimSpace(historyScanner.Text())

			// fileTail, err := tail.TailFile(filePath, tail.Config{
			// 	Follow: false, // Do not wait for new lines
			// 	ReOpen: false, // Do not reopen the file if it changes
			// })
			// if err != nil {
			// 	if os.IsNotExist(err) {
			// 		fmt.Printf("Skipping %s, file does not exist.\n", filePath)
			// 		continue
			// 	}
			// 	panic(err)
			// }

			lines := strings.SplitSeq(string(line), "\n")
			for line := range lines {

				if line == "" {
					continue
				}

				fullLine := strings.Fields(line)
				if len(fullLine) == 0 {
					continue
				}

				var commandTime time.Time
				if strings.HasPrefix(line, ":") && strings.Contains(line, ";") {
					parts := strings.SplitN(line, ";", 2)

					line = parts[1]
					firstPart := strings.SplitN(parts[0], ":", 3)

					strUnixTimeWithoutSpace := strings.TrimSpace(firstPart[1])
					intUnixTime, err := strconv.ParseInt(strUnixTimeWithoutSpace, 10, 32)
					if err != nil {
						log.Fatal(err)
					}

					commandTime = time.Unix(intUnixTime, 0)
					line = parts[1]

					if time.Now().Sub(commandTime) <= 24*time.Hour && time.Now().After(commandTime) {
						tui.OtherCount["Commands today"]++
					}

				}

				fields := strings.Fields(line)

				if len(fields) == 0 {
					continue
				}

				var option string
				cmd := fields[0]
				if len(fields) > 1 {

					option = fields[1]

					if cmd == "sudo" {
						cmd = fields[1]
					}
					if option == "commit" && len(fields) > 2 {
						tui.CommitCount["Commits Overall"]++
						if commandTime.After(time.Now().Add(-24 * time.Hour)) {
							tui.CommitCount["Commits Today"]++
						}
					}
				} else {
					continue
				}
				tui.CommandCount[cmd]++
			}
		}
		defer file.Close()
	}
	tui.CreateTextView()
}
