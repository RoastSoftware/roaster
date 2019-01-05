// Package model feed implementation.
package model

import (
	"time"
)

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

// GetFeed collects the latest N (N = pageSize) feed items for either everyone,
// followees or for a user.
//
// * An empty username string returns the feed items for all users.
// * A set username together with the followees parameter set to true returns the
//   users followees feed items.
// * A set username together with the followees parameter set to false returns
//   only the users feed items.
//
// Pagination is supported where page = 0 is the first (latest) page.
func GetFeed(username string, followees bool, page uint64, pageSize uint64) (feed Feed, err error) {
	rows, err := database.Query(`
	SELECT DISTINCT r.username, r.score, r.language, r.create_time
	FROM "roaster"."roast" AS r
		LEFT OUTER JOIN "roaster"."user_followees" AS f
		ON $2 AND LOWER(f.username)=LOWER(TRIM($1))
	WHERE
		COALESCE(TRIM($1), '')='' OR
		NOT $2 AND LOWER(r.username)=LOWER(TRIM($1)) OR
		$2 AND r.username = f.followee
	ORDER BY r.create_time DESC LIMIT $3 OFFSET $4
	`, username, followees, pageSize, pageSize*page)
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
