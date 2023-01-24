package internals

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Generate(historyFile string, word string) {

	readFile, err := os.Open(historyFile)

	if err != nil {
		ThrowError(err)
	}

	tempFile, err := os.Create("/tmp/history.txt")

	if err != nil {
		ThrowError(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		split := strings.Split(line, ";")

		commands := split[len(split)-1]

		count := strings.Count(commands, word)

		if count == 0 {

			_, err := tempFile.WriteString(line + "\n")

			if err != nil {
				ThrowError(err)
			}

		}

	}

	defer readFile.Close()
	defer tempFile.Close()

}

func Clear(historyFile string) {

	err := os.Remove(historyFile)

	if err != nil {
		ThrowError(err)
	}

	oldLocation := "/tmp/history.txt"
	newLocation := historyFile

	copyerr := copy(oldLocation, newLocation)

	if copyerr != nil {
		ThrowError(copyerr)
	}

}

func copy(src, dst string) error {

	_, err := os.Stat(src)
	if err != nil {
		ThrowError(err)
	}

	source, err := os.Open(src)
	if err != nil {
		ThrowError(err)
	}

	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		ThrowError(err)
	}

	defer destination.Close()

	_, copyerr := io.Copy(destination, source)

	return copyerr
}
