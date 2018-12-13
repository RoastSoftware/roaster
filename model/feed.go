package model

import (
	"time"
)

const PageSize = 25

const (
	RoastCategory = iota
)

type FeedItem struct {
	Category    int       `json:"category"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"createTime"`
}

type Feed struct {
	Items []FeedItem `json:"items"`
}

// GetGlobalFeed collects the latest N (N = PageSize) feed items for all users.
// Pagination is supported where page = 0 is the first (latest) page.
func GetGlobalFeed(page uint64) (feed Feed, err error) {
	return getFeed("", page)
}

// GetUserFeed collects the latest N (N = PageSize) feed items for an user.
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
	`, username, PageSize, PageSize*page)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := FeedItem{}
		err = rows.Scan(&item.Username, &item.Title, &item.Description, &item.CreateTime)
		if err != nil {
			return Feed{}, err
		}

		feed.Items = append(feed.Items, item)
	}

	return
}
