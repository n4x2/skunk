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

	fmt.Printf("enter vault password: ")
	secret, err := terminal.AskCredentials()
	if err != nil {
		return fmt.Errorf("\n%w", err)
	}

	if secret == "" {
		return fmt.Errorf("\n%w", &EmptyValueError{Field: "vault password"})
	}

	passwords, err := pass.ListPassword(secret)
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if passwords == nil {
		fmt.Printf("\nvault: no password available\n")
		return nil
	}

	fmt.Printf("\n\navailable %d passwords:\n", len(passwords))
	for _, password := range passwords {
		fmt.Printf("- %s\n", password.Name)
	}
	return nil
}
