package data

import "time"

var LineFields []string
var CommandTime time.Time

type CommandUses struct {
	Command string
	Uses    int
}

type Commits struct {
	CommitType string
	Commits    int
}
