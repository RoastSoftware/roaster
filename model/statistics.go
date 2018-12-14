// Package model feed implementation.
package model

import (
	"time"
)

// LanguageDatapoint represents a datapoint for language statistics.
type LanguageDatapoint struct {
	Start    time.Time
	End      time.Time
	Errors   uint64
	Warnings uint64
	Rows     uint64
}

// LanguageStatistics holds a collection of datapoints.
type LanguageStatistics struct {
	Datapoints []LanguageDatapoint `json:"items"`
}

// GetGlobalLinesOfCodeForLanguage
func GetGlobalLinesOfCodeForLanguage(language string) (lines uint64, err error) {
	return getGlobalLinesOfCode(language)
}

// GetGlobalLinesOfCode
func GetGlobalLinesOfCode() (lines uint64, err error) {
	return getGlobalLinesOfCode("")
}

// GetUserLinesOfCodeForLanguage
func GetUserLinesOfCodeForLanguage(username, language string) (lines uint64, err error) {
	return getUserLinesOfCode(username, language)
}

// GetUserLinesOfCode
func GetUserLinesOfCode(username string) (lines uint64, err error) {
	return getUserLinesOfCode(username, "")
}

// GetGlobalNumberOfRoasts
func GetGlobalNumberOfRoasts() (numberOfRoasts uint64, err error) {
	return getNumberOfRoasts("")
}

// GetUserNumberOfRoasts
func GetUserNumberOfRoasts(username string) (numberOfRoasts uint64, err error) {
	return getNumberOfRoasts(username)
}

// GetUserScore
func GetUserScore(username string) (score uint64, err error) {
	err = database.QueryRow(`
		SELECT SUM(score)
		FROM roaster.roast AS r
		WHERE LOWER(r.username)=LOWER(TRIM($1))
	`, username).Scan(&score)
	if err != nil {
		return
	}

	return
}

func getGlobalLinesOfCode(language string) (lines uint64, err error) {
	err = database.QueryRow(`
		SELECT SUM(lines_of_code)
		FROM roaster.roast_statistics AS s
		JOIN roaster.roast AS r
		ON r.id = s.roast
		WHERE COALESCE(TRIM($2), '')='' OR
		      LOWER(r.language)=LOWER(TRIM($2))
	`, language).Scan(&lines)
	if err != nil {
		return
	}

	return
}

func getUserLinesOfCode(username, language string) (lines uint64, err error) {
	err = database.QueryRow(`
		SELECT SUM(lines_of_code)
		FROM roaster.roast_statistics AS s
		JOIN roaster.roast AS r
		ON r.id = s.roast
		WHERE LOWER(r.username)=LOWER(TRIM($1)) AND
		      (COALESCE(TRIM($2), '')='' OR
		       LOWER(r.language)=LOWER(TRIM($2)))
	`, username, language).Scan(&lines)
	if err != nil {
		return
	}

	return
}

func getNumberOfRoasts(username string) (numberOfRoasts uint64, err error) {
	err = database.QueryRow(`
		SELECT COUNT(username)
		FROM roaster.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(username)=LOWER(TRIM($1))
	`, username).Scan(&numberOfRoasts)
	if err != nil {
		return
	}

	return
}
