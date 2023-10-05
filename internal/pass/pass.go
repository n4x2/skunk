// Package pass provides functions for managing password data in a vault.
package pass

import (
	"encoding/json"
	"strings"

	"github.com/n4x2/skunk/internal/vault"
)

// Password represents a structure for storing password information.
type Password struct {
	Name     string
	Username string
	Pass     string
	Token    string
}

// getVault retrieves the vault and passwords.
func getVault(pass string) (string, []Password, error) {
	var passwords []Password

	vaultFile, exists := vault.IsVaultExist()
	if !exists {
		// Create a new vault if it doesn't exist.
		newVaultFile, err := vault.NewVault(pass, passwords)
		if err != nil {
			return "", nil, err
		}
		return newVaultFile, passwords, nil
	}

	decoded, err := vault.Decode(vaultFile, pass)
	if err != nil {
		return "", nil, err
	}

	if err := json.Unmarshal(decoded, &passwords); err != nil {
		return "", nil, err
	}

	return vaultFile, passwords, nil
}

// AddPassword adds a new password to the vault.
func AddPassword(value Password, pass string) error {
	vaultFile, passwords, err := getVault(pass)
	if err != nil {
		return err
	}

	passwords = append(passwords, value)
	if err := vault.Encode(vaultFile, pass, passwords); err != nil {
		return err
	}

	return nil
}

// FindPassword retrive [Password] by matching the name.
func FindPassword(name, pass string) (Password, error) {
	var password Password

	_, passwords, err := getVault(pass)
	if err != nil {
		return password, err
	}

	for _, pass := range passwords {
		if strings.EqualFold(pass.Name, name) {
			return pass, nil
		}
	}

	return password, nil
}

// RemovePassword remove a password from vault.
func RemovePassword(name, pass string) error {
	vaultFile, passwords, err := getVault(pass)
	if err != nil {
		return err
	}

	var newPasswords []Password
	for _, password := range passwords {
		if !strings.EqualFold(password.Name, name) {
			newPasswords = append(newPasswords, password)
		}
	}

	if err := vault.Encode(vaultFile, pass, newPasswords); err != nil {
		return err
	}

	return nil
}
