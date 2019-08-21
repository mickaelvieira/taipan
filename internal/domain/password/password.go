package password

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
)

// IsValid is the password strong enough
func IsValid(p string) bool {
	return len(p) >= 10
}

// Hash generates a hash for the given password
func Hash(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "password hashing")
	}
	return string(h), nil
}
