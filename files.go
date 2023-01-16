package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Generate(historyFile string) {

	readFile, err := os.Open(historyFile)

	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	tempFile, err := os.Create("/tmp/history.txt")

	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		split := strings.Split(line, ";")

		commands := split[len(split)-1]

		count := strings.Count(commands, "clear")

		if count == 0 {

			_, err := tempFile.WriteString(line + "\n")

			if err != nil {
				fmt.Printf("[ERROR] : %v\n", err.Error())
			}

		}

	}

	defer readFile.Close()
	defer tempFile.Close()

}

func Clear(historyFile string) {

	err := os.Remove(historyFile)

	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	oldLocation := "/tmp/history.txt"
	newLocation := historyFile

	copyerr := copy(oldLocation, newLocation)

	if copyerr != nil {
		fmt.Printf("[ERROR] : %v\n", copyerr.Error())
	}

}

func copy(src, dst string) error {

	_, err := os.Stat(src)
	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	source, err := os.Open(src)
	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	defer destination.Close()

	_, copyerr := io.Copy(destination, source)

	return copyerr
}
