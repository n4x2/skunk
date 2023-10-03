// Package vault utilizes [age] package for secure data encryption and decryption.
//
// [age]: https://pkg.go.dev/filippo.io/age
package vault

const (
	Perm600 = 0o600 // For secret files.
	Perm700 = 0o700 // For sensitive files or directories.
)
