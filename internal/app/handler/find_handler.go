package handler

import (
	"flag"
	"fmt"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
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
			return &EmptyValueError{Field: "field: name"}
		}
	}

	fmt.Printf("enter vault password: ")
	secret, err := terminal.AskCredentials()
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if secret == "" {
		return &EmptyValueError{Field: "vault password"}
	}

	password, err := pass.FindPassword(name, secret)
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if password.Name == "" {
		return fmt.Errorf("\nerror: password \"%s\" is not found", name)
	}

	fmt.Printf("\n\"%s\" is available\n", password.Name)
	return nil
}
