package handler

import (
	"flag"
	"fmt"
	"os"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
	"golang.org/x/term"
)

func ListPassword(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Vault password.
	fmt.Printf("Enter vault password: ")
	passVault, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	passwords, err := pass.ListPassword(string(passVault))
	if err != nil {
		terminal.ClearLines(1)
		return err
	}

	if passwords == nil {
		terminal.ClearLines(1)
		fmt.Println("No password available.")
		return nil
	}

	terminal.ClearLines(1)
	fmt.Printf("Available passwords:\n\n")
	for _, password := range passwords {
		fmt.Println(password.Name)
	}

	return nil
}
