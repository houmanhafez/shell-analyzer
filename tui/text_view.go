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

func CreateTextView(app *tview.Application) *tview.Flex {
	top := tview.NewTextView()
	left := tview.NewTextView()
	rightTop := tview.NewTextView()
	rightBottom := tview.NewTextView()
	bottom := tview.NewTextView()

	go func() {
		var sortedTopCmdValues []data.TopCmds
		var sortedGitCmdDailyValues []data.GitCmds
		var sortedSystemCmdDailyValues []data.SystemCmds
		var sortedGitCmdTotalValues []data.GitCmds
		var sortedSystemCmdTotalValues []data.SystemCmds

		for command, uses := range UnsortedTopCmds {
			sortedTopCmdValues = append(sortedTopCmdValues, data.TopCmds{Command: command, Uses: uses})
		}

		for commitType, commits := range UnsortedGitCmdsDaily {
			sortedGitCmdDailyValues = append(sortedGitCmdDailyValues, data.GitCmds{Command: commitType, Uses: commits})
		}

		for commitType, commits := range UnsortedGitCmdsTotal {
			sortedGitCmdTotalValues = append(sortedGitCmdTotalValues, data.GitCmds{Command: commitType, Uses: commits})
		}

		for command, uses := range UnsortedSystemCmdsTotal {
			sortedSystemCmdTotalValues = append(sortedSystemCmdTotalValues, data.SystemCmds{Command: command, Uses: uses})
		}
		for command, uses := range UnsortedSystemCmdsDaily {
			sortedSystemCmdDailyValues = append(sortedSystemCmdDailyValues, data.SystemCmds{Command: command, Uses: uses})
		}

		sort.Slice(sortedTopCmdValues, func(i, j int) bool {
			return sortedTopCmdValues[i].Uses > sortedTopCmdValues[j].Uses
		})

		for i, frame := range data.ProgressBar {
			currentFrame := i
			app.QueueUpdateDraw(func() {
				top.SetTextAlign(tview.AlignCenter).SetText("\n" + frame + "\n")
				if currentFrame == len(data.ProgressBar)-1 {
					top.SetText("")
				}
			})
			time.Sleep(30 * time.Millisecond)
		}

		app.QueueUpdateDraw(func() {
			fmt.Fprintf(top, "\n[white::b]DOS BIOS (2A4FVCXL)\nSHELL ANALYZING UTILITY - COPYRIGHT (C) 1985-2005, AMERICAN MEGATRON INC.\n\n")
		})

		app.QueueUpdateDraw(func() {
			left.SetTextAlign(tview.AlignLeft).SetText("\n     [red::b]TOP 20 COMMANDS YOU'VE USED[-]\n\n")
			for i, kv := range sortedTopCmdValues {
				if i >= 20 {
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

				fmt.Fprintf(left, "     [%s]%2d. %-30s - %5d times\n", color, i+1, kv.Command, kv.Uses)
			}
		})

		app.QueueUpdateDraw(func() {
			fmt.Fprintf(rightTop, "\n     [yellow::b]REPORT OF TODAY\n\n")

			for _, kv := range sortedGitCmdDailyValues {
				fmt.Fprintf(rightTop, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
			}

			for _, kv := range sortedSystemCmdDailyValues {
				fmt.Fprintf(rightTop, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
			}
		})

		app.QueueUpdateDraw(func() {
			fmt.Fprintf(rightBottom, "\n     [yellow::b]OVERALL USAGE\n\n")

			for _, kv := range sortedGitCmdTotalValues {
				fmt.Fprintf(rightBottom, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
			}
			for _, kv := range sortedSystemCmdTotalValues {
				fmt.Fprintf(rightBottom, "     [white]%-34s - %5d times\n", kv.Command, kv.Uses)
			}
		})

		app.QueueUpdateDraw(func() {
			fmt.Fprintf(bottom, "\nTo Be Continued\n\n")
		})
	}()

	top.SetDynamicColors(true).SetWrap(true).SetBackgroundColor(tcell.NewHexColor(0x0000AA))
	rightBottom.SetDynamicColors(true).SetWrap(true).SetBackgroundColor(tcell.NewHexColor(0x0000AA)).SetBorder(true)
	rightTop.SetDynamicColors(true).SetWrap(true).SetBackgroundColor(tcell.NewHexColor(0x0000AA)).SetBorder(true)
	left.SetDynamicColors(true).SetWrap(true).SetBackgroundColor(tcell.NewHexColor(0x0000AA)).SetBorder(true)
	bottom.SetTextAlign(tview.AlignCenter).SetDynamicColors(true).SetWrap(true).SetBackgroundColor(tcell.NewHexColor(0x0000AA))

	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(rightTop, 0, 1, false).
		AddItem(rightBottom, 0, 1, false)

	centerFlex := tview.NewFlex().
		AddItem(left, 0, 1, false).
		AddItem(rightFlex, 0, 1, false)

	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(top, 3, 0, false).
		AddItem(centerFlex, 0, 1, true).
		AddItem(bottom, 3, 0, false)

	mainFlex.SetBorder(true).SetBackgroundColor(tcell.ColorWhite)

	return mainFlex
}
