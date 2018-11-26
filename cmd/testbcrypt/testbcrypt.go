// Package testbcrypt tests the Go Bcrypt implementation for common problems.
package testbcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashWithSamePassword generates a bcrypt hash for the same password with different cost input.
func HashWithSamePassword(cost int) []byte {
	res, _ := bcrypt.GenerateFromPassword([]byte("G;H.#Wj9PLH<>TmkgzDn{?FY&U_"), cost)
	return res
}

// HashWithSameCost generates a bcrypt hash with the default cost.
func HashWithSameCost(password string) []byte {
	res, _ := bcrypt.GenerateFromPassword([]byte(password), -1)
	return res
}
