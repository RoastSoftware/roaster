// Package model feed implementation.
package model

import (
	"time"
)

const pageSize = 25

const (
	roastCategory = iota
)

// FeedItem represents a item in a feed.
type FeedItem struct {
	Category    int       `json:"category"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"createTime"`
}

// Feed holds a list of FeedItems.
type Feed struct {
	Items []FeedItem `json:"items"`
}

// GetGlobalFeed collects the latest N (N = pageSize) feed items for all users.
// Pagination is supported where page = 0 is the first (latest) page.
func GetGlobalFeed(page uint64) (feed Feed, err error) {
	return getFeed("", false, page)
}

// GetUserFeed collects the latest N (N = pageSize) feed items for an user.
// Pagination is supported where page = 0 is the first (latest) page.
func GetUserFeed(username string, friends bool, page uint64) (feed Feed, err error) {
	return getFeed(username, friends, page)
}

// TODO: Add friends to SQL query when the Friends table has been implemented.
func getFeed(username string, friends bool, page uint64) (feed Feed, err error) {
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
