package main

import (
	"fmt"

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

func main() {
	res1 := HashWithSameCost("abc\x00def")
	if err := bcrypt.CompareHashAndPassword(res1, []byte("abc\x00ghi")); err != nil {
		fmt.Println("Go does _not_ have the null byte bug 👍")
	} else {
		fmt.Println("Go has the null byte bug 👎")
	}

	s72b := "104751087632048762130750394869807019873409852374095283740952837423452345"
	res1 = HashWithSameCost(s72b + "A")
	if err := bcrypt.CompareHashAndPassword(res1, []byte(s72b+"B")); err != nil {
		fmt.Println("Go does _not_ truncate passwords after 72 bytes 👍")
	} else {
		fmt.Println("Go truncates passwords after 72 bytes 👎")
	}
}
