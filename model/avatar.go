// Package model avatar implementation.
package model

import (
	"bytes"
	"database/sql"
	"image"
	"image/png"
	"math/rand"
	"time"

	"github.com/o1egl/govatar"
)

// Avatar holds a avatar.
type Avatar struct {
	Username string
	Avatar   []byte
}

// PutAvatar adds a avatar to a user in the database.
func PutAvatar(avatar Avatar) (err error) {
	_, err = database.Exec(`
        INSERT INTO "avatar"
        (username, avatar)
        VALUES
        (TRIM($1), $2)
        ON CONFLICT
        (username)
        DO UPDATE SET 
        avatar = EXCLUDED.avatar
    `, avatar.Username, avatar.Avatar)
	return
}

// EncodeToPNG encodes a Go image.Image to a PNG.
func encodeToPNG(img image.Image) (raw []byte, err error) {
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	raw = buf.Bytes()
	return
}

// GetAvatar return a avatar by their username.
// Note: Username will be trimmed for whitespace before quering the database.
func GetAvatar(username string) (avatar Avatar, err error) {
	avatar, err = getAvatar(username)
	if err == sql.ErrNoRows {
		err = nil
		var img image.Image
		var raw []byte

		img, err = govatar.GenerateFromUsername(randomGender(), username)
		if err != nil {
			return
		}

		raw, err = encodeToPNG(img)
		if err != nil {
			return
		}

		avatar = Avatar{username, raw}
		err = PutAvatar(avatar)
	}

	return
}

func getAvatar(username string) (avatar Avatar, err error) {
	err = database.QueryRow(`
        SELECT username, avatar
        FROM "avatar"
        WHERE username=TRIM($1)
        `, username).Scan(&avatar.Username, &avatar.Avatar)
	return
}

func randomGender() govatar.Gender {
	r := rand.Intn(2)
	if r == 0 {
		return govatar.MALE
	}
	return govatar.FEMALE
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
