// Package terminal provides utilities for terminal-related tasks.
package terminal

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// AskConfirmation returns true if it's "y" or "Y", otherwise false.
func AskConfirmation() bool {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			return true
		}
	}

	return false
}

// AskValue reads from stdin and returns it as a string.
func AskValue() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

// AskCredentials reads a password from the terminal without echoing it,
// returns the entered password as a string.
func AskCredentials() (string, error) {
	secret, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	return string(secret), nil
}

// ClearLines clears specified lines in terminal.
func ClearLines(num int) {
	for i := 0; i < num; i++ {
		fmt.Print("\033[2K\r")
		if num == 1 {
			// Move the cursor to left.
			fmt.Print("\033[D")
			break
		}
		fmt.Print("\033[A")
	}
}
