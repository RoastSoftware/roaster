// Package model user implementation.
package model

import (
	"crypto/sha512"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Set Bcrypt work factor so it takes >250 ms on a modern CPU.
const bcryptWorkFactor = 12

// User holds a user.
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

// generateHash hashes the password as:
// 	hash = bcrypt(sha512(password))
//
// Bcrypt is used as the core hashing algorithm with 12 rounds.
//
// The plaintext password that is provided will first be transformed to a
// hash sum with SHA-512. This is due to that Bcrypt limits the input to 72
// bytes. By hashing the password with SHA-512 more entropy of the original
// password is kept. Also, some implementations of Bcrypt that allows for longer
// passwords can be vulnerable to DoS attacks[0].
//
// The SHA-512 hash sum is then hashed again using Bcrypt. This is because
// SHA-512 is a _fast_ hash algorithm not made for password hashing. Bcrypt is
// designed to be slow and hard to speed up using hardware such as FPGAs and
// ASICs. The work factor is set to 12 which should make the expensive Blowfish
// setup take >250 ms (364.815906 ms precisely on my sucky laptop).
//
// Dropbox has a great article[1] on their password hashing scheme which our
// scheme shares many similarities with, we do not, however, use AES-256 with a
// global pepper (shared global encryption key) which is overkill for our use
// case. They also encode their SHA-512 with base 64, which is not needed in our
// case.
//
// Some implementations of Bcrypt uses a null byte (\x00) to determine the end
// of the input, the Go implementations does _not_ have this problem. So there
// is no need to encode the data as base 64. This is verified using the program
// in cmd/testbcrypt.
//
// The approach used by Dropbox where they encode with base 64 will generate a
// ~88 byte long key, which is then truncated to 72 bytes. This results in a
// input with 64^72 possible combinations. Our approach of not encoding the
// input as base 64 results in a 64 byte long key (the output size of SHA-512),
// where there is 256 possible combinations _per byte_. This results in a input
// with 256^64 possible combinations. Therefore our approach allows for far more
// possible entropy because 64^72 << 256^64 (like 99.999... % more entropy ;)).
//
// [0]: https://arstechnica.com/information-technology/2013/09/long-passwords-are-good-but-too-much-length-can-be-bad-for-security/
// [1]: https://blogs.dropbox.com/tech/2016/09/how-dropbox-securely-stores-your-passwords/
func generateHash(password []byte) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(sha512Sum(password), -1)

	return
}

// sha512Sum wraps the sha512.Sum512 method and returns a variadic byte slice
// instead of the fixed size byte array that sha512.Sum512 provides.
//
// Note: The function will also zero out any variables that is passed (except
// for the actual output).
func sha512Sum(in []byte) (out []byte) {
	sum := sha512.Sum512(in)
	in = []byte{}              // Empty the password variable.
	out = sum[:]               // Convert to byte slice of variadic length.
	sum = [sha512.Size]byte{0} // Empty the sum variable.

	return out
}

// validPassword validates a hash against a provided password.
// See the generateHash function for documentation about the algorithms used.
func validPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, sha512Sum(password))
	if err != nil {
		return false
	}

	return true
}

// PutUser adds a user to the database. The password parameter is a plain-text
// password that will be hashed before insertion. Note that the password
// parameter value will be replaced with a empty slice.
//
// Note: The username and e-mail for the user will be trimmed from whitespace.
func PutUser(user User, password []byte) (err error) {
	var hash []byte
	defer func() {
		hash = []byte{} // Empty hash on function return.
	}()

	hash, err = generateHash(password)
	if err != nil {
		return
	}

	_, err = database.Exec(`
		INSERT INTO "roaster"."user"
		(username, email, fullname, create_time, hash)
		VALUES
		(TRIM($1), LOWER(TRIM($2)), $3, $4, $5)
	`, user.Username, user.Email, user.Fullname, time.Now(), hash)

	return
}

// AuthenticateUser authenticates the user with their username/e-mail and
// password.
//
// The identifier parameter can be the users username or e-mail address.
//
// If the authentication fails an empty user with the ok boolean set to false is
// returned. Else the user struct is filled and the ok boolean is set to true.
//
// Note: That the password parameter will be emptied.
func AuthenticateUser(identifier string, password []byte) (user User, ok bool) {
	var hash []byte
	defer func() {
		hash = []byte{} // Empty hash on function return.
	}()

	err := database.QueryRow(`
		SELECT username, email, fullname, hash
		FROM "roaster"."user"
		WHERE (LOWER(username)=LOWER(TRIM($1)) OR email=LOWER(TRIM($1)))
	`, identifier).Scan(&user.Username, &user.Email, &user.Fullname, &hash)
	if err != nil {
		ok = false
		return
	}

	// The parameters hash and password are replaced with an empty slice by
	// the validPassword function.
	ok = validPassword(hash, password)
	return
}

// GetUser returns a user by their username or e-mail.
//
// Note: The username and e-mail for the user will be trimmed from whitespace
// before searching in the database.
func GetUser(identifier string) (user User, err error) {
	err = database.QueryRow(`
		SELECT username, email, fullname
		FROM "roaster"."user"
		WHERE (LOWER(username)=LOWER(TRIM($1)) OR email=LOWER(TRIM($1)))
	`, identifier).Scan(&user.Username, &user.Email, &user.Fullname)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = fmt.Errorf(`user: "%s" does not exist`, identifier)
		default:
			log.Println(err)
			err = errors.New("failed to execute your request")
		}
	}

	return
}
