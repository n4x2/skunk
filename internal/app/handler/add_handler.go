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
			return fmt.Errorf("value can not empty")
		}
	}

	// Username.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("(%s) Username: ", name)
	if scanner.Scan() {
		username = scanner.Text()
	}

	// Password.
	fmt.Printf("(%s) Password: ", name)
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	password = string(passwordBytes)

	// OTP.
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\nAre you want add OTP for account "+`"%s"`+" ? (y/n): ", username)
	if scanner.Scan() {
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			scanner = bufio.NewScanner(os.Stdin)
			fmt.Printf("\n(OTP) Token: ")
			if scanner.Scan() {
				token = scanner.Text()
			}
		}
	}

	// Vault password.
	fmt.Printf("\nEnter vault password: ")
	passVault, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	input := pass.Password{
		Name:     name,
		Username: username,
		Pass:     password,
		Token:    token,
	}

	if err = pass.AddPassword(input, string(passVault)); err != nil {
		terminal.ClearLines(5)
		return err
	}

	terminal.ClearLines(5)
	fmt.Printf(`"%s" successfully saved.`+"\n", name)

	return nil
}
