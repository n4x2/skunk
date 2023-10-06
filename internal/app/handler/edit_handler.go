package handler

import (
	"flag"
	"fmt"

	"github.com/n4x2/skunk/internal/pass"
	"github.com/n4x2/skunk/internal/terminal"
)

// EditPassword edit existing password in vault.
func EditPassword(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		fs.Usage()
		return nil
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	var name string

	var (
		passname, otp, password, username bool
	)

	if nameFlag := fs.Lookup("name"); nameFlag != nil {
		name = nameFlag.Value.String()
		if name == "" {
			return &EmptyValueError{Field: "flag: name"}
		}
	}

	fmt.Print("enter vault password: ")
	secret, err := terminal.AskCredentials()
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if secret == "" {
		return fmt.Errorf("\n%w", &EmptyValueError{Field: "password"})
	}

	oldPassword, err := pass.FindPassword(name, secret)
	if err != nil {
		return fmt.Errorf("\nerror: %w", err)
	}

	if oldPassword.Name == "" {
		return fmt.Errorf("\n\"%s\" not found", name)
	}

	name = oldPassword.Name

	fmt.Printf("\n(%s) edit password name ? (y/n): ", name)
	if ok := terminal.AskConfirmation(); ok {
		fmt.Print("new name: ")
		if value := terminal.AskValue(); value != "" {
			oldPassword.Name = value
			passname = true
		}
	}

	fmt.Printf("(%s) edit username ? (y/n): ", name)
	if ok := terminal.AskConfirmation(); ok {
		fmt.Printf("new username: ")
		if value := terminal.AskValue(); value != "" {
			oldPassword.Username = value
			username = true
		}
	}

	fmt.Printf("(%s) update password ? (y/n): ", name)
	if ok := terminal.AskConfirmation(); ok {
		fmt.Printf("new password: ")
		secret, err := terminal.AskCredentials()
		if err != nil {
			return fmt.Errorf("\nerror: %w", err)
		}
		oldPassword.Pass = secret
		password = true
		fmt.Print("\n")
	}

	fmt.Printf("(%s) edit OTP ? (y/n): ", name)
	if ok := terminal.AskConfirmation(); ok {
		fmt.Printf("new token: ")
		if value := terminal.AskValue(); value != "" {
			oldPassword.Token = value
			otp = true
		}
	}

	if password || otp || username || passname {
		fmt.Printf("\nconfirm to update \"%s\" ? (y/n): ", name)
		if ok := terminal.AskConfirmation(); ok {
			if err := pass.UpdatePassword(oldPassword, name, secret); err != nil {
				return err
			}
			fmt.Printf("\n\"%s\" successfully updated\n", name)
			return nil
		}
	}

	fmt.Println("\nnothing change were made, edit has been canceled")
	return nil
}
