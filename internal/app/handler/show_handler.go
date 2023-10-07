package handler

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
	"github.com/n4x2/zoo/to"
)

// ShowPassword show existing password by name.
func ShowPassword(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		fs.Usage()
		return nil
	}

	var (
		copy         bool
		name, secret string
	)

	if nameFlag := fs.Lookup("name"); nameFlag != nil {
		name = nameFlag.Value.String()
		if name == "" {
			return &EmptyValueError{Field: "flag: name"}
		}
	}

	if copyFlag := fs.Lookup("copy"); copyFlag != nil {
		value, err := to.Bool(copyFlag.Value.String())
		if err != nil {
			return fmt.Errorf("\nerror: %w", err)
		}
		copy = value
	}

	passFlag := fs.Lookup("pass")
	if passFlag != nil {
		secret = passFlag.Value.String()
	}

	if secret == "" {
		fmt.Print("enter vault password: ")
		value, err := terminal.AskCredentials()
		if err != nil {
			return fmt.Errorf("\nerror: %w", err)
		}
		secret = value
	}

	password, err := pass.FindPassword(name, secret)
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if password.Name == "" {
		fmt.Printf("\nvault: \"%s\" not available\n", name)
		return nil
	}

	if copy {
		err := clipboard.WriteAll(password.Pass)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}

		fmt.Printf("\n\npassword \"%s\" copied into clipboard\n", name)
		return nil
	}

	fmt.Printf("\n\n\"%s\"\n", password.Pass)
	return nil
}
