package handler

import (
	"flag"
	"fmt"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
)

func ListPassword(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}

	var secret string
	passFlag := fs.Lookup("pass")
	if passFlag != nil {
		secret = passFlag.Value.String()
	}

	if secret == "" {
		fmt.Printf("enter vault password: ")
		value, err := terminal.AskCredentials()
		if err != nil {
			return fmt.Errorf("\n%w", err)
		}

		if value == "" {
			return fmt.Errorf("\n%w", &EmptyValueError{Field: "vault password"})
		}
		secret = value
		fmt.Println()
	}

	passwords, err := pass.ListPassword(secret)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	if passwords == nil {
		fmt.Printf("\nvault: no password available\n")
		return nil
	}

	for _, password := range passwords {
		fmt.Printf("- %s\n", password.Name)
	}
	return nil
}
