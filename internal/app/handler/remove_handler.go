package handler

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
	"golang.org/x/term"
)

// RemovePassword remove password from vault by name.
func RemovePassword(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		fs.Usage()
		return nil
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	var name string
	if nameFlag := fs.Lookup("name"); nameFlag != nil {
		name = nameFlag.Value.String()
		if name == "" {
			return fmt.Errorf("value can not empty")
		}
	}

	// Vault password.
	fmt.Printf("Enter vault password: ")
	passVault, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	password, err := pass.FindPassword(name, string(passVault))
	if err != nil {
		terminal.ClearLines(1)
		return err
	}

	if password.Name == "" {
		terminal.ClearLines(1)
		return fmt.Errorf(`"%s" not found`, name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nAre you sure want to remove "+`"%s" from vault`+" ? (y/n): ", password.Name)
	if scanner.Scan() {
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			err := pass.RemovePassword(name, string(passVault))
			if err != nil {
				terminal.ClearLines(2)
				return err
			}

			terminal.ClearLines(2)
			fmt.Printf(`"%s" successfully removed`+"\n", password.Name)
		}
	}

	return nil
}
