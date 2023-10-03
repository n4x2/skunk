// Package passgen wraps package [pass] to generate random passwords.
//
// [pass]: https://pkg.go.dev/github.com/n4x2/zoo/pass
package passgen

import "github.com/n4x2/zoo/pass"

// GeneratePassword generates a random password with options. It can
// include numbers if `number` is true and symbols if `symbol` is true.
// The length of the password is determined by the `len` parameter.
// It returns error if `len` <= 0.
func GeneratePassword(number, symbol bool, len int) (string, error) {
	pass, err := pass.Generate(number, symbol, len)
	if err != nil {
		return "", err
	}

	return pass, nil
}
