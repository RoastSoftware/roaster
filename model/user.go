// Package model user implementation.
package model

import (
	"crypto/sha512"
	"fmt"
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

// UserScore holds a named uint64 that represents the users score.
type UserScore struct {
	Score uint64 `json:"score"`
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
func sha512Sum(in []byte) []byte {
	sum := sha512.Sum512(in)
	in = []byte{} // Empty the password variable.
	return sum[:]
}

// validPassword validates a hash against a provided password.
// See the generateHash function for documentation about the algorithms used.
func validPassword(hash, password []byte) bool {
	sum := sha512Sum(password)
	defer func() {
		sum = []byte{}
	}()

	err := bcrypt.CompareHashAndPassword(hash, sum)
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
		($1, TRIM($2), TRIM($3), $4, $5)
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
	return
}

// UpdateUser updates fullname, password or email of a user, otherwise error.
func UpdateUser(user User) (err error) {
	var modifiedUsername string

	err = database.QueryRow(`
    UPDATE "roaster"."user" SET 
        fullname = COALESCE(NULLIF($2, ''), fullname),
        email = COALESCE(NULLIF($3, ''), email)
    WHERE LOWER(username)=LOWER(TRIM($1))
    RETURNING username
    `, user.Username, user.Fullname, user.Email).Scan(&modifiedUsername)

	if err != nil {
		return
	}

	if modifiedUsername != user.Username {
		err = fmt.Errorf("failed to update provided username (%s != %s)",
			modifiedUsername,
			user.Username)
	}

	return
}

// GetUserScore returns the score for an user.
func GetUserScore(username string) (score UserScore, err error) {
	err = database.QueryRow(`
		SELECT COALESCE(SUM("score"), 0) AS "score"
		FROM roaster.roast AS r
		WHERE LOWER(r.username)=LOWER(TRIM($1))
	`, username).Scan(&score.Score)
	return
}

// Follower represents a user that is tracking someone.
type Follower struct {
	Username   string    `json:"username"`
	CreateTime time.Time `json:"createTime"`
}

// Followee represents a user that is being tracked by someone.
type Followee Follower

// GetFollowees returns a list of followees if successful, otherwise error.
func GetFollowees(username string) (followees []Followee, err error) {
	followeeRows, err := database.Query(`
    SELECT followee, create_time
    FROM "roaster"."user_followees"
    WHERE (LOWER(username)=LOWER(TRIM($1)))
    `, username)
	if err != nil {
		return
	}
	defer followeeRows.Close()

	for followeeRows.Next() {
		res := Followee{}
		err = followeeRows.Scan(&res.Username, &res.CreateTime)
		if err != nil {
			return
		}
		followees = append(followees, res)
	}
	return
}

// PutFollowee saves a followee relation to the DB, returns error if unsuccessful.
func PutFollowee(username string, followee string) (err error) {
	_, err = database.Exec(`
    INSERT INTO "roaster"."user_followees"
    (username, create_time, followee)
    VALUES
    (TRIM($1), $2, TRIM($3))
    `, username, time.Now(), followee)
	return
}

// RemoveFollowee deletes a followee relation from DB, returns error if unsuccessful.
func RemoveFollowee(username string, followee string) (err error) {
	_, err = database.Exec(`
    DELETE FROM "roaster"."user_followees"
    WHERE (lower(username)=lower(TRIM($1)) AND lower(followee)=lower(TRIM($2)))
    `, username, followee)
	return
}

// GetFollowers gets a list of followers for a specific user, returns error if unsuccsessful.
func GetFollowers(username string) (followers []Follower, err error) {
	followerRows, err := database.Query(`
    SELECT username, create_time
    FROM "roaster"."user_followees"
    WHERE (LOWER(followee)=LOWER(TRIM($1)))
    `, username)
	if err != nil {
		return
	}
	defer followerRows.Close()

	for followerRows.Next() {
		res := Follower{}
		err = followerRows.Scan(&res.Username, &res.CreateTime)
		if err != nil {
			return
		}
		followers = append(followers, res)
	}
	return
}
