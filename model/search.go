// Package model search implementation.
package model

import (
	"fmt"
	"time"
)

const userCategory = iota

// SearchResult holds a search result.
type SearchResult struct {
	Category    int    `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// SearchUsers searchs for all users matching the provided query.
func SearchUsers(query string) (results []SearchResult, err error) {
	rows, err := database.Query(`
	SELECT username, fullname, create_time FROM roaster.user
		-- Naive search, can be extended with Full Text Search in the
		-- future.
		WHERE username LIKE $1 OR
	   	      fullname LIKE $1 OR
	   	      email LIKE $1
	ORDER BY username DESC
	`, fmt.Sprintf("%%%s%%", query))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username, fullname string
		var createTime time.Time

		err = rows.Scan(&username, &fullname, &createTime)
		if err != nil {
			return []SearchResult{}, err
		}

		result := SearchResult{
			Category:    userCategory,
			Title:       username,
			Description: fmt.Sprintf("User %s, also known as %s. Member since %s.", username, fullname, createTime.Format("02 January 2006")),
			URL:         fmt.Sprintf("/user/%s", username),
		}

		results = append(results, result)
	}

	return
}

// SearchAll returns a list of SearchResults.
func SearchAll(query string) (results []SearchResult, err error) {
	// Only searching for users is supported, but in the future several
	// search call results can be merged into one slice by this function.
	return SearchUsers(query)
}
