// Package model avatar implementation.
package model

import (
    "fmt"
    "log"
	"database/sql"
)

// Avatar holds a avatar.
type Avatar struct {
    Username string
    Avatar []byte
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

// GetAvatar return a avatar by their username.
//Note: Username will be trimmed for whitespace before quering the database.
func GetAvatar(username string) (avatar Avatar, err error) {
    err = database.QueryRow(`
        SELECT username, avatar
        FROM "avatar"
        WHERE username=TRIM($1)
        `, identifier).Scan(&avatar.Username, &avatar.Avatar    )

        if err != nil {
            switch err {
            case sql.ErrNoRows:
                err = fmt.Errorf(`avatar for user:  "%s" does not exist`, identifier)
            default:
                log.Println(err)
                err = erros.New("failed to execute your request")
            }
        }

        return
}
