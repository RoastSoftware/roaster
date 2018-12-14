// Package model feed implementation.
package model

import (
	"time"
)

const pageSize = 25

const (
	roastCategory = iota
)

// LanguageDatapoint represents a datapoint for language statistics.
type LanguageDatapoint struct {
	Start    time.Time
	End      time.Time
	Errors   uint64
	Warnings uint64
	Rows     uint64
}

// LanguageStatistics holds a
type LanguageStatistics struct {
	Datapoints []LanguageDatapoint `json:"items"`
}

// GetGlobalFeed collects the latest N (N = pageSize) feed items for all users.
// Pagination is supported where page = 0 is the first (latest) page.
func GetGlobalFeed(page uint64) (feed Feed, err error) {
	return getFeed("", page)
}

// GetUserFeed collects the latest N (N = pageSize) feed items for an user.
// Pagination is supported where page = 0 is the first (latest) page.
func GetUserFeed(username string, page uint64) (feed Feed, err error) {
	return getFeed(username, page)
}

func getFeed(username string, page uint64) (feed Feed, err error) {
	rows, err := database.Query(`
	SELECT username, score, language, create_time
	FROM "roaster"."roast"
	WHERE coalesce(TRIM($1), '')='' OR LOWER(username)=LOWER(TRIM($1))
	ORDER BY create_time DESC LIMIT $2 OFFSET $3
	`, username, pageSize, pageSize*page)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		// NOTE: Hardcoded category as only the Roast category is
		// supported.
		item := FeedItem{Category: roastCategory}

		err = rows.Scan(&item.Username, &item.Title, &item.Description, &item.CreateTime)
		if err != nil {
			return Feed{}, err
		}

		feed.Items = append(feed.Items, item)
	}

	return
}
