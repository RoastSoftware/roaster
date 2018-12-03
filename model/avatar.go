// Package model avatar implementation.
package model

import (
    "fmt"
    "log"
	"database/sql"
)

type Avatar struct {
    Username string
    Avatar []byte
}


