package data

import "os"

var Files = []string{
	os.Getenv("HOME") + "/.bash_history",
	os.Getenv("HOME") + "/.zsh_history",
}
