package tui

import (
	"fmt"
	"shell-analyzer/m/data"
	"sort"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var UnsortedTopCmds = make(map[string]int)
var UnsortedGitCmds = make(map[string]int)
var UnsortedSystemCmds = make(map[string]int)

func CreateTextView() *tview.TextView {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	go func() {

		var sortedTopCmdValues []data.TopCmds
		var sortedGitCmdValues []data.GitCmds
		var sortedSystemCmdValues []data.TopCmds

		for command, uses := range UnsortedTopCmds {
			sortedTopCmdValues = append(sortedTopCmdValues, data.TopCmds{Command: command, Uses: uses})
		}

		for commitType, commits := range UnsortedGitCmds {
			sortedGitCmdValues = append(sortedGitCmdValues, data.GitCmds{Command: commitType, Uses: commits})
		}

		for command, uses := range UnsortedSystemCmds {
			sortedSystemCmdValues = append(sortedSystemCmdValues, data.TopCmds{Command: command, Uses: uses})
		}

		sort.Slice(sortedTopCmdValues, func(i, j int) bool {
			return sortedTopCmdValues[i].Uses > sortedTopCmdValues[j].Uses
		})

		for i, frame := range data.ProgressBar {
			app.QueueUpdateDraw(func() {

				textView.SetTextAlign(tview.AlignCenter).SetText("\n" + frame + "\n")
				if i == len(data.ProgressBar)-1 {
					textView.SetText("")
				}
			})

			time.Sleep(30 * time.Millisecond)
		}

		textView.SetTextAlign(tview.AlignLeft).SetText("\n     [::b][yellow]Top 10 commands you've used\n\n")

		for i, kv := range sortedTopCmdValues {
			if i >= 10 {
				break
			}
			fmt.Fprintf(textView, "     [white]%d %-10s - %d times\n", i+1, kv.Command, kv.Uses)
		}

		fmt.Fprintf(textView, "\n\n     [yellow]Other Facts\n\n")

		for _, kv := range sortedGitCmdValues {
			fmt.Fprintf(textView, "     [white]%-s - %d\n", kv.Command, kv.Uses)
		}

		for _, kv := range sortedSystemCmdValues {
			fmt.Fprintf(textView, "\n     [white]%-s - %d\n", kv.Command, kv.Uses)
		}
	}()

	textView.SetDoneFunc(func(key tcell.Key) { app.Stop() })
	textView.SetDynamicColors(true).SetWrap(true)
	textView.SetBackgroundColor(tcell.NewHexColor(0x0000AA))
	textView.SetTitle("Shell Analyzer").SetTitleColor(tcell.ColorBlack)
	textView.SetBorder(true).SetBackgroundColor(tcell.ColorWhite)

	if err := app.SetRoot(textView, true).SetFocus(textView).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	return textView
}
