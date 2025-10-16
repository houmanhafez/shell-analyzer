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

func NewTextView() *tview.TextView {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
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

		var sortedCommandUses []data.CommandUses
		var sortedCommitValues []data.Commits

		for command, uses := range CommandCount {
			sortedCommandUses = append(sortedCommandUses, data.CommandUses{Command: command, Uses: uses})
		}

		for commitType, commits := range CommitCount {
			sortedCommitValues = append(sortedCommitValues, data.Commits{CommitType: commitType, Commits: commits})
		}

		sort.Slice(sortedCommandUses, func(i, j int) bool {
			return sortedCommandUses[i].Uses > sortedCommandUses[j].Uses
		})

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
		for i, kv := range sortedCommandUses {
			if i >= 10 {
				break
			}
			fmt.Fprintf(textView, "     [white]%d %-10s - %d times\n", i+1, kv.Command, kv.Uses)
		}
		fmt.Fprintf(textView, "\n\n     Other Facts\n\n")

		for _, kv := range sortedCommitValues {

			fmt.Fprintf(textView, "     %-s - %d times\n", kv.CommitType, kv.Commits)
		}
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
	return textView
}
