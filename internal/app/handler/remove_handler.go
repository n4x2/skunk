package handler

import (
	"flag"
	"fmt"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
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
			return &EmptyValueError{Field: "flag: name"}
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

	fmt.Printf("\n\nconfirm removal of \"%s\" from vault? (y/n): ", password.Name)
	if ok := terminal.AskConfirmation(); ok {
		err := pass.RemovePassword(name, secret)
		if err != nil {
			return fmt.Errorf("\nerror: %w", err)
		}

		fmt.Printf("\n\"%s\" has been successfully removed\n", password.Name)
		return nil
	}

	fmt.Println("\nno changes were made, password deletion has been canceled")
	return nil
}
