// Package model feed implementation.
package model

import (
	"fmt"
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

// RoastCountDatapoint represents a single datapoint with a timestamp and the
// count of Roasts during that timestamp.
type RoastCountDatapoint struct {
	Timestamp time.Time `json:"timestamp"`
	Count     uint64    `json:"count"`
}

// RoastCountTimeseries holds a slice of RoastCountDatapoints.
type RoastCountTimeseries []RoastCountDatapoint

// GetGlobalLinesOfCodeForLanguage returns the number of lines of code for all
// users for a specific language.
func GetGlobalLinesOfCodeForLanguage(language string) (lines uint64, err error) {
	return getGlobalLinesOfCode(language)
}

// GetGlobalLinesOfCode returns the number of lines of code for all users for
// all languages.
func GetGlobalLinesOfCode() (lines uint64, err error) {
	return getGlobalLinesOfCode("")
}

// GetUserLinesOfCodeForLanguage returns the number of lines of code for an user
// for a specific language.
func GetUserLinesOfCodeForLanguage(username, language string) (lines uint64, err error) {
	return getUserLinesOfCode(username, language)
}

// GetUserLinesOfCode returns the number of lines of code for an user for all
// languages.
func GetUserLinesOfCode(username string) (lines uint64, err error) {
	return getUserLinesOfCode(username, "")
}

// GetGlobalNumberOfRoasts returns the number of Roasts for all users and
// languages.
func GetGlobalNumberOfRoasts() (numberOfRoasts uint64, err error) {
	return getNumberOfRoasts("")
}

// GetUserNumberOfRoasts returns the number of Roasts for a specific user.
func GetUserNumberOfRoasts(username string) (numberOfRoasts uint64, err error) {
	return getNumberOfRoasts(username)
}

// GetUserScore returns the score for an user.
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

// GetGlobalRoastCountTimeseries returns a timeseries between a start and end
// timestamp, the interval parameter should be any duration above 1 minute,
// anything less will default to 1 minute.
func GetGlobalRoastCountTimeseries(start, end time.Time, resolution time.Duration) (
	timeseries RoastCountTimeseries, err error) {

	return getRoastCountTimeseries(start, end, resolution, "")
}

// GetUserRoastCountTimeseries returns a timeseries between a start and end
// timestamp for the provided username, the interval parameter should be any
// duration above 1 minute, anything less will default to 1 minute.
func GetUserRoastCountTimeseries(start, end time.Time, resolution time.Duration, username string) (
	timeseries RoastCountTimeseries, err error) {

	return getRoastCountTimeseries(start, end, resolution, username)
}

// getRoastCountTimeseries returns a timeseries of number of Roasts per time
// unit. See: GetGlobalRoastCountTimeseries or GetUserRoastCountTimeseries.
//
// The minimum interval is 1 minute, anything less will be set to 1 minute per
// default.
func getRoastCountTimeseries(start, end time.Time, interval time.Duration, username string) (
	timeseries RoastCountTimeseries, err error) {

	// Round the interval to the closest minute.
	interval = interval.Round(time.Minute)
	if interval < 1 { // 1 minute is the minimum interval.
		// Just enforce 1 minute interval instead of erroring.
		interval = 1 * time.Minute
	}

	// Round both the start and end time to the nearest multiple of the
	// interval.
	start = start.Round(interval)
	end = end.Round(interval)

	rows, err := database.Query(`
		WITH "time_series" AS (
			SELECT generate_series(
				$1::timestamp, -- Start point date with time.
				$2::timestamp, -- End point date with time.
				$3::interval   -- Time series interval.
			) AS "datapoint"
		)

		SELECT
			"time_series"."datapoint" AS "timestamp",
			COUNT(r."id") AS "count"
		FROM "time_series"
		LEFT JOIN "roaster"."roast" AS r
			-- Truncate and compare per resolution.
			ON date_trunc('minute', "roaster".round_minutes(r."create_time"::timestamp, $4)) = "time_series"."datapoint"
			AND (COALESCE(TRIM($5), '')='' OR LOWER(r."username")=LOWER(TRIM($5))) -- Optionally only return for specified user.
		GROUP BY 1
		ORDER BY "timestamp" DESC -- First in result is the latest timestamp.
		LIMIT 1000 -- Do not return more than 1000 rows back in time.
	`,
		start,
		end,
		fmt.Sprintf("%.0f minutes", interval.Minutes()),
		int(interval.Minutes()),
		username)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		datapoint := RoastCountDatapoint{}

		err = rows.Scan(&datapoint.Timestamp, &datapoint.Count)
		if err != nil {
			return
		}

		timeseries = append(timeseries, datapoint)
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
