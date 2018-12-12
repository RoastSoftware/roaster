package model

import "time"

const PageSize = 25

type Feed struct {
	Category    int
	Username    string
	Title       string
	Description string
	CreateTime  time.Time
}

func GetGlobalFeed(page uint) (feed Feed, err error) {
	start := page * PageSize
	end := start + PageSize

	rows, err := database.Query(`
	-- TODO
	`, start, end)

	rows.Next()

	return
}
