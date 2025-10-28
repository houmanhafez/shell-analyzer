package data

import "time"

var LineFields []string
var CmdTime time.Time

type TopCmds struct {
	Command string
	Uses    int
}

type GitCmds struct {
	Command string
	Uses    int
}

type SystemCmds struct {
	Command string
	Uses    int
}
