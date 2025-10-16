package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type keyValue struct {
	Key   string
	Value int
}

var files = []string{
	os.Getenv("HOME") + "/.bash_history",
	os.Getenv("HOME") + "/.zsh_history",
}

func main() {

	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	commandCounts := make(map[string]int)
	commitCount := make(map[string]int)

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
						commitCount["Commits so far"]++
					}
				} else {
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
	var sortedCommitValues []keyValue

	for key, value := range commandCounts {
		sortedKeyValues = append(sortedKeyValues, keyValue{key, value})
	}

	for key, value := range commitCount {
		sortedCommitValues = append(sortedCommitValues, keyValue{key, value})
	}

	sort.Slice(sortedKeyValues, func(i, j int) bool {
		return sortedKeyValues[i].Value > sortedKeyValues[j].Value
	})

	levels := []string{
		"[white] ░░░░░░░░░░░░░░",
		"[red]   █░░░░░░░░░░░░░",
		"[red]   ██░░░░░░░░░░░░",
		"[red]   ███░░░░░░░░░░░",
		"[red]   ████░░░░░░░░░░",
		"[red]   █████░░░░░░░░░",
		"[yellow]██████░░░░░░░░",
		"[yellow]███████░░░░░░░",
		"[yellow]████████░░░░░░",
		"[yellow]█████████░░░░░",
		"[yellow]██████████░░░░",
		"[green] ███████████░░░",
		"[green] ████████████░░",
		"[green] █████████████░",
		"[green] ██████████████",
	}

	go func() {
		for i, frame := range levels {
			app.QueueUpdateDraw(func() {

				textView.SetTextAlign(tview.AlignCenter).SetText("\n" + frame + "\n")
				if i == len(levels)-1 {
					textView.SetText("")
				}
			})

			time.Sleep(100 * time.Millisecond)
		}

		textView.SetTextAlign(tview.AlignLeft).SetText("\n     [::b][red]Top 10 commands you've used\n\n")
		for i, kv := range sortedKeyValues {
			if i >= 10 {
				break
			}
			fmt.Fprintf(textView, "     [white]%d %-10s - %d times\n", i+1, kv.Key, kv.Value)
		}
		fmt.Fprintf(textView, "\n\n     Other Facts\n\n")

		for _, kv := range sortedCommitValues {

			fmt.Fprintf(textView, "     %-s - %d times\n", kv.Key, kv.Value)
		}
	}()

	go func() {

	}()

	textView.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})
	textView.SetDynamicColors(true).SetWrap(true)

	textView.SetBackgroundColor(tcell.NewHexColor(0x0000AA))

	textView.SetTitle("Shell Analyzer").SetTitleColor(tcell.ColorBlack)

	textView.SetBorder(true).SetBackgroundColor(tcell.ColorWhite)

	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}
}
