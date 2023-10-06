package handler

import (
	"flag"
	"fmt"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
)

// AddPassword add new password based on flags.
func AddPassword(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		fs.Usage()
		return nil
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	var (
		name, username, password, token string
	)

	if nameFlag := fs.Lookup("name"); nameFlag != nil {
		name = nameFlag.Value.String()
		if name == "" {
			return &EmptyValueError{Field: "flag: name"}
		}
	}

	fmt.Printf("(%s) username: ", name)
	value := terminal.AskValue()
	if value == "" {
		return &EmptyValueError{Field: "username"}
	}
	username = value

	fmt.Printf("(%s) password: ", name)
	secret, err := terminal.AskCredentials()
	if err != nil {
		return fmt.Errorf("\n%w", err)
	}

	if secret == "" {
		return fmt.Errorf("\n%w", &EmptyValueError{Field: "password"})
	}
	password = secret

	fmt.Printf("\n\nadd OTP for \"%s\" ? (y/n): ", username)
	if ok := terminal.AskConfirmation(); ok {
		fmt.Printf("\n(OTP) Token: ")
		value := terminal.AskValue()
		if value == "" {
			return &EmptyValueError{Field: "OTP token"}
		}
		token = value
	}

	fmt.Printf("\nEnter vault password: ")
	secret, err = terminal.AskCredentials()
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if secret == "" {
		return fmt.Errorf("\n%w", &EmptyValueError{Field: "vault password"})
	}

	input := pass.Password{
		Name:     name,
		Username: username,
		Pass:     password,
		Token:    token,
	}

	if err = pass.AddPassword(input, secret); err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	fmt.Printf("\n\"%s\" successfully saved\n", name)
	return nil
}
