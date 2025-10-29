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

var UnsortedGitCmdsDaily = make(map[string]int)
var UnsortedSystemCmdsDaily = make(map[string]int)

var UnsortedGitCmdsTotal = make(map[string]int)
var UnsortedSystemCmdsTotal = make(map[string]int)

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

		var sortedGitCmdDailyValues []data.GitCmds
		var sortedSystemCmdDailyValues []data.TopCmds

		var sortedGitCmdTotalValues []data.GitCmds
		var sortedSystemCmdTotalValues []data.TopCmds

		for command, uses := range UnsortedTopCmds {
			sortedTopCmdValues = append(sortedTopCmdValues, data.TopCmds{Command: command, Uses: uses})
		}

		for commitType, commits := range UnsortedGitCmdsDaily {
			sortedGitCmdDailyValues = append(sortedGitCmdDailyValues, data.GitCmds{Command: commitType, Uses: commits})
		}

		for command, uses := range UnsortedSystemCmdsTotal {
			sortedSystemCmdDailyValues = append(sortedSystemCmdDailyValues, data.TopCmds{Command: command, Uses: uses})
		}

		for commitType, commits := range UnsortedGitCmdsTotal {
			sortedGitCmdTotalValues = append(sortedGitCmdTotalValues, data.GitCmds{Command: commitType, Uses: commits})
		}

		for command, uses := range UnsortedSystemCmdsTotal {
			sortedSystemCmdTotalValues = append(sortedSystemCmdTotalValues, data.TopCmds{Command: command, Uses: uses})
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

		textView.SetTextAlign(tview.AlignLeft).SetText("\n     [red::b]Top 10 commands you've used[-]\n\n")
		for i, kv := range sortedTopCmdValues {
			if i >= 10 {
				break
			}

			var color string
			switch i {
			case 0:
				color = "gold"
			case 1:
				color = "silver"
			case 2:
				color = "orange"
			default:
				color = "white"
			}

			fmt.Fprintf(textView, "     [%s]%2d. %-30s - %5d times\n", color, i+1, kv.Command, kv.Uses)
		}

		fmt.Fprintf(textView, "\n\n     [yellow]Report of Today\n\n")

		for _, kv := range sortedGitCmdDailyValues {
			fmt.Fprintf(textView, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
		}

		for _, kv := range sortedSystemCmdDailyValues {
			fmt.Fprintf(textView, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
		}

		fmt.Fprintf(textView, "\n\n     [yellow]Overall\n\n")

		for _, kv := range sortedGitCmdTotalValues {
			fmt.Fprintf(textView, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
		}
		for _, kv := range sortedSystemCmdTotalValues {
			fmt.Fprintf(textView, "\n     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
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
