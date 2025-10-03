package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type keyValue struct {
	Key   string
	Value int
}

var files = []string{
	os.Getenv("HOME") + "/.bash_history",
	os.Getenv("HOME") + "/.zsh_history",
	// os.Getenv("HOME") + "/.local/share/fish/fish_history",
}

func main() {

	commandCounts := make(map[string]int)

	for _, filePath := range files {
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

			lines := strings.Split(string(line), "\n")

			for _, line := range lines {
				firstWord := strings.Fields(line)
				if len(firstWord) == 0 {
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

				cmd := fields[0]
				if cmd == "sudo" && len(fields) > 1 {
					cmd = fields[1]
				}

				commandCounts[cmd]++
			}
		}
		defer file.Close()
	}
	var sortedKeyValues []keyValue

	for key, value := range commandCounts {
		sortedKeyValues = append(sortedKeyValues, keyValue{key, value})
	}

	sort.Slice(sortedKeyValues, func(i, j int) bool {
		return sortedKeyValues[i].Value > sortedKeyValues[j].Value
	})

	for i, kv := range sortedKeyValues {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %-10s - %d times\n", i+1, kv.Key, kv.Value)
	}
}
