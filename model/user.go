// Package model user implementation.
package model

// User holds a user.
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}
