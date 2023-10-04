// Package terminal provides utilities for terminal-related tasks.
package terminal

import "fmt"

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
