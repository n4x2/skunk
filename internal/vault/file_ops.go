package vault

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"filippo.io/age"
	"filippo.io/age/armor"
)

// Decode decrypts the specified file using the given passphrase.
// It returns the decrypted data as a byte slice. It returns error
// if the decryption or file operations fail.
func Decode(file, pass string) ([]byte, error) {
	rawFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()

	identity, err := age.NewScryptIdentity(pass)
	if err != nil {
		return nil, err
	}

	armorFile := armor.NewReader(rawFile)

	decryptedFile, err := age.Decrypt(armorFile, identity)
	if err != nil {
		return nil, err
	}

	value := &bytes.Buffer{}

	_, err = io.Copy(value, decryptedFile)
	if err != nil {
		return nil, err
	}

	return value.Bytes(), nil
}

// Encode encrypts given value and writes it to the specified file
// using the given passphrase. It returns an error if the encryption
// or file operations fail.
func Encode(file, pass string, value interface{}) error {
	tmpFile, err := os.CreateTemp(filepath.Dir(file), "*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if err = tmpFile.Chmod(Perm600); err != nil {
		return err
	}

	recipient, err := age.NewScryptRecipient(pass)
	if err != nil {
		return err
	}

	armorFile := armor.NewWriter(tmpFile)

	encryptedFile, err := age.Encrypt(armorFile, recipient)
	if err != nil {
		return err
	}

	if err = json.NewEncoder(encryptedFile).Encode(value); err != nil {
		return err
	}

	if err = encryptedFile.Close(); err != nil {
		return err
	}

	if err = armorFile.Close(); err != nil {
		return err
	}

	if err = tmpFile.Sync(); err != nil {
		return err
	}

	if err = tmpFile.Close(); err != nil {
		return err
	}

	if err = os.Rename(tmpFile.Name(), file); err != nil {
		return err
	}

	return nil
}
