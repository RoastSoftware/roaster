// Package model user implementation.
package model

import (
	"crypto/sha512"
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
// Bcrypt is used as the core hashing algorithm with 10 rounds.
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
// case.
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
		INSERT INTO user
		(username, email, fullname, create_time, hash)
		VALUES
		(TRIM(?), LOWER(TRIM(?)), ?, ?)
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

	err := database.QueryRow(`
		SELECT username, email, fullname, hash
		FROM user
		WHERE (username=TRIM(?) OR email=LOWER(TRIM(?)))
	`, identifier).Scan(&user.Email, &user.Username, &user.Fullname, &hash)
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
		FROM user
		WHERE (username=TRIM(?) OR email=LOWER(TRIM(?)))
	`, identifier).Scan(&user.Email, &user.Username, &user.Fullname)

	return
}
