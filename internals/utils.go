package internals

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func GetUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}
	return currentUser.Username
}

func GetShell() string {
	file_path := "/etc/passwd"
	currentUser, err := user.Current()
	var currentShell string

	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	username := currentUser.Username

	readFile, err := os.Open(file_path)

	if err != nil {
		fmt.Printf("[ERROR] : %v\n", err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		arr := strings.Split(line, ":")

		if arr[0] == username {
			shellPath := arr[len(arr)-1]
			split := strings.Split(shellPath, "/")

			currentShell = split[2]
		}

	}

	defer readFile.Close()

	return currentShell
}
