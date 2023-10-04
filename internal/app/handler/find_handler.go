package handler

import (
	"flag"
	"fmt"
	"os"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
	"golang.org/x/term"
)

// FindPassword find password by matching the name.
func FindPassword(fs *flag.FlagSet, args []string) error {
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
	if err != nil || password.Name == "" {
		terminal.ClearLines(1)
		return fmt.Errorf(`Password "%s" `+"not found", name)
	}

	terminal.ClearLines(1)
	fmt.Printf(`"%s" is available`+"\n", password.Name)

	return nil
}
