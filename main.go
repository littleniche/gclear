package main

import (
	"time"

	"github.com/briandowns/spinner"
)

type Shell string

const (
	bash Shell = "/.bash_history"
	zsh        = "/.zsh_history"
	fish       = "/.local/share/fish/fish_history"
)

const (
	RedText    = "\033[31m"
	NormalText = "\033[0m"
)

func main() {

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Clearing history..."

	currentShell := GetShell()

	var historyFile Shell

	switch currentShell {
	case "bash":
		historyFile = bash
	case "zsh":
		historyFile = zsh
	}

	username := GetUsername()

	path := "/home/" + username + string(historyFile)

	s.Start()
	time.Sleep(500 * time.Millisecond)
	Generate(path)
	s.Stop()

	Clear(path)

}
