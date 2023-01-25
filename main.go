package main

import (
	"flag"
	"time"

	"github.com/briandowns/spinner"
	"github.com/littleniche/gclear/internals"
)

type Shell string

const (
	bash Shell = "/.bash_history"
	zsh        = "/.zsh_history"
	fish       = "/.local/share/fish/fish_history"
)


func main() {

	var word string = "clear"

	notClear := flag.Bool("w", false, "delete a specific word")
	flag.Parse()

	if *notClear {
		word = flag.Arg(0)
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Clearing history...\n"

	currentShell := internals.GetShell()

	var historyFile Shell

	switch currentShell {
	case "bash":
		historyFile = bash
	case "zsh":
		historyFile = zsh
	case "fish":
		historyFile = fish
	}

	username := internals.GetUsername()

	path := "/home/" + username + string(historyFile)

	s.Start()
	time.Sleep(500 * time.Millisecond)
	internals.Generate(path, word);
	s.Stop()

	internals.Clear(path)

}
