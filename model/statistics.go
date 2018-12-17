// Package model feed implementation.
package model

import (
	"fmt"
	"time"
)

// RoastDatapoint represents a single datapoint with a timestamp.
type RoastDatapoint struct {
	Timestamp        time.Time `json:"timestamp"`
	Count            uint64    `json:"count"`
	NumberOfErrors   uint64    `json:"numberOfErrors"`
	NumberOfWarnings uint64    `json:"numberOfWarnings"`
	LinesOfCode      uint64    `json:"linesOfCode"`
}

// RoastTimeseries holds a slice of RoastDatapoints.
type RoastTimeseries []RoastDatapoint

// NumberOfRoasts holds a named uint64 that represents the number of Roasts.
type NumberOfRoasts struct {
	Count uint64 `json:"count"`
}

// LinesOfCode holds a named uint64 that represents the number of lines.
type LinesOfCode struct {
	Lines uint64 `json:"lines"`
}

// GetGlobalLinesOfCodeForLanguage returns the number of lines of code for all
// users for a specific language.
func GetGlobalLinesOfCodeForLanguage(language string) (lines LinesOfCode, err error) {
	return getGlobalLinesOfCode(language)
}

// GetGlobalLinesOfCode returns the number of lines of code for all users for
// all languages.
func GetGlobalLinesOfCode() (lines LinesOfCode, err error) {
	return getGlobalLinesOfCode("")
}

// GetUserLinesOfCodeForLanguage returns the number of lines of code for an user
// for a specific language.
func GetUserLinesOfCodeForLanguage(username, language string) (lines LinesOfCode, err error) {
	return getUserLinesOfCode(username, language)
}

// GetUserLinesOfCode returns the number of lines of code for an user for all
// languages.
func GetUserLinesOfCode(username string) (lines LinesOfCode, err error) {
	return getUserLinesOfCode(username, "")
}

// GetGlobalNumberOfRoasts returns the number of Roasts for all users and
// languages.
func GetGlobalNumberOfRoasts() (numberOfRoasts NumberOfRoasts, err error) {
	return getNumberOfRoasts("")
}

// GetUserNumberOfRoasts returns the number of Roasts for a specific user.
func GetUserNumberOfRoasts(username string) (numberOfRoasts NumberOfRoasts, err error) {
	return getNumberOfRoasts(username)
}

// GetGlobalRoastTimeseries returns a timeseries between a start and end
// timestamp, the interval parameter should be any duration above 1 minute,
// anything less will default to 1 minute.
func GetGlobalRoastTimeseries(start, end time.Time, resolution time.Duration) (
	timeseries RoastTimeseries, err error) {

	return getRoastTimeseries(start, end, resolution, "")
}

// GetUserRoastTimeseries returns a timeseries between a start and end
// timestamp for the provided username, the interval parameter should be any
// duration above 1 minute, anything less will default to 1 minute.
func GetUserRoastTimeseries(start, end time.Time, resolution time.Duration, username string) (
	timeseries RoastTimeseries, err error) {

	return getRoastTimeseries(start, end, resolution, username)
}

// getRoastTimeseries returns a timeseries of number of Roasts per time
// unit. See: GetGlobalRoastTimeseries or GetUserRoastTimeseries.
//
// The minimum interval is 1 minute, anything less will be set to 1 minute per
// default.
func getRoastTimeseries(start, end time.Time, interval time.Duration, username string) (
	timeseries RoastTimeseries, err error) {

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
			COUNT(r."id") AS "count",
			COALESCE(SUM(s."number_of_errors"), 0) AS "number_of_errors",
			COALESCE(SUM(s."number_of_warnings"), 0) AS "number_of_warnings",
			COALESCE(SUM(s."lines_of_code"), 0) AS "lines_of_code"
		FROM "time_series"
		LEFT JOIN "roaster"."roast" AS r
			-- Truncate and compare per resolution.
			ON date_trunc('minute', "roaster".round_minutes(r."create_time"::timestamp, $4)) = "time_series"."datapoint"
			AND (COALESCE(TRIM($5), '')='' OR LOWER(r."username")=LOWER(TRIM($5))) -- Optionally only return for specified user.
			LEFT JOIN "roaster"."roast_statistics" AS s -- Collect statistics for the Roasts.
				ON s."roast" = r."id"
		GROUP BY 1
		ORDER BY "timestamp" DESC -- First in result is the latest timestamp.
		LIMIT 1000 -- Do not return more than 1000 rows back in time.
	`,
		start,
		end,
		fmt.Sprintf("%.0f minutes", interval.Minutes()),
		interval.Minutes(),
		username)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		datapoint := RoastDatapoint{}

		err = rows.Scan(
			&datapoint.Timestamp,
			&datapoint.Count,
			&datapoint.NumberOfErrors,
			&datapoint.NumberOfWarnings,
			&datapoint.LinesOfCode)
		if err != nil {
			return
		}

		timeseries = append(timeseries, datapoint)
	}

	return
}

func getGlobalLinesOfCode(language string) (lines LinesOfCode, err error) {
	err = database.QueryRow(`
		SELECT SUM(lines_of_code)
		FROM roaster.roast_statistics AS s
		JOIN roaster.roast AS r
		ON r.id = s.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(r.language)=LOWER(TRIM($1))
	`, language).Scan(&lines.Lines)
	if err != nil {
		return
	}

	return
}

func getUserLinesOfCode(username, language string) (lines LinesOfCode, err error) {
	err = database.QueryRow(`
		SELECT SUM(lines_of_code)
		FROM roaster.roast_statistics AS s
		JOIN roaster.roast AS r
		ON r.id = s.roast
		WHERE LOWER(r.username)=LOWER(TRIM($1)) AND
		      (COALESCE(TRIM($2), '')='' OR
		       LOWER(r.language)=LOWER(TRIM($2)))
	`, username, language).Scan(&lines.Lines)
	if err != nil {
		return
	}

	return
}

func getNumberOfRoasts(username string) (numberOfRoasts NumberOfRoasts, err error) {
	err = database.QueryRow(`
		SELECT COUNT(username)
		FROM roaster.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(username)=LOWER(TRIM($1))
	`, username).Scan(&numberOfRoasts.Count)
	if err != nil {
		return
	}

	return
}
