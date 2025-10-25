package tui

import (
	"fmt"
	"shell-analyzer/m/data"
	"sort"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var CommandCount = make(map[string]int)
var CommitCount = make(map[string]int)
var OtherCount = make(map[string]int)

func CreateTextView() *tview.TextView {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	go func() {

		var sortedCommandUses []data.CommandUses
		var sortedCommitValues []data.Commits
		var sortedOtherValues []data.CommandUses

		for command, uses := range CommandCount {
			sortedCommandUses = append(sortedCommandUses, data.CommandUses{Command: command, Uses: uses})
		}

		for commitType, commits := range CommitCount {
			sortedCommitValues = append(sortedCommitValues, data.Commits{CommitType: commitType, Commits: commits})
		}

		for command, uses := range OtherCount {
			sortedOtherValues = append(sortedOtherValues, data.CommandUses{Command: command, Uses: uses})
		}

		sort.Slice(sortedCommandUses, func(i, j int) bool {
			return sortedCommandUses[i].Uses > sortedCommandUses[j].Uses
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
		for i, kv := range sortedCommandUses {
			if i >= 10 {
				break
			}
			fmt.Fprintf(textView, "     [white]%d %-10s - %d times\n", i+1, kv.Command, kv.Uses)
		}
		fmt.Fprintf(textView, "\n\n     [yellow]Other Facts\n\n")

		for _, kv := range sortedCommitValues {
			fmt.Fprintf(textView, "     [white]%-s - %d\n", kv.CommitType, kv.Commits)
		}

		for _, kv := range sortedOtherValues {
			fmt.Fprintf(textView, "     [white]%-s - %d\n", kv.Command, kv.Uses)
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
