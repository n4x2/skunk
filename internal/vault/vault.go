// Package vault utilizes [age] package for secure data encryption and decryption.
//
// [age]: https://pkg.go.dev/filippo.io/age
package vault

import (
	"os"
	"path/filepath"
)

const (
	Perm600 = 0o600 // For secret files.
	Perm700 = 0o700 // For sensitive files or directories.
)

var (
	File = "vault"  // Default vault file name.
	Dir  = ".skunk" // Default vault directory.
)

// getVaultFilePath returns vault file path.
func getVaultFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, Dir, File), nil
}

// IsVaultExist checks if the vault file exists and returns the path.
func IsVaultExist() (string, bool) {
	vaultFile, err := getVaultFilePath()
	if err != nil {
		return "", false
	}

	_, err = os.Stat(vaultFile)
	return vaultFile, !os.IsNotExist(err)
}

// NewVault creates a new vault file with the provided password and
// initial value. It return the vault file path.
func NewVault(pass string, value interface{}) (string, error) {
	vaultFile, err := getVaultFilePath()
	if err != nil {
		return "", err
	}

	vaultDir := filepath.Dir(vaultFile)

	if _, err := os.Stat(vaultDir); os.IsNotExist(err) {
		if err := os.MkdirAll(vaultDir, Perm700); err != nil {
			return "", err
		}
	}

	if err := Encode(vaultFile, pass, value); err != nil {
		return "", err
	}

	return vaultFile, nil
}
