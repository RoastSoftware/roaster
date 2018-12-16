// Package model feed implementation.
package model

import (
	"fmt"
	"time"
)

const (
	// SecondResolution represents a single second.
	SecondResolution = time.Second
	// MinuteResolution is exactly 60 seconds.
	MinuteResolution = time.Minute
	// HourResolution is exactly 60 minutes.
	HourResolution = time.Hour
	// DayResolution is exactly 24 hours.
	DayResolution = time.Hour * 24
	// MonthResolution is exactly 30 days.
	MonthResolution = DayResolution * 30
	// YearResolution is exactly 365 days.
	YearResolution = DayResolution * 365
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
// timestamp, the resolution parameters should be any of:
//  * SecondResolution,
//  * MinuteResolution,
//  * HourResolution,
//  * DayResolution   - exactly 24 hours,
//  * MonthResolution - exactly 30 days,
//  * YearResolution  - exactly 365 days.
//
// If the resolution doesn't match any of the supported resolutions above, it'll
// default to the YearResolution.
//
// The resolutions are exact and does not take into account leap seconds etc.
func GetGlobalRoastCountTimeseries(start, end time.Time, resolution time.Duration) (
	timeseries RoastCountTimeseries, err error) {

	return getRoastCountTimeseries(start, end, resolution, "")
}

// GetUserRoastCountTimeseries returns a timeseries between a start and end
// timestamp for the provided username, the resolution parameters should be any
// of:
//  * SecondResolution,
//  * MinuteResolution,
//  * HourResolution,
//  * DayResolution   - exactly 24 hours,
//  * MonthResolution - exactly 30 days,
//  * YearResolution  - exactly 365 days.
//
// If the resolution doesn't match any of the supported resolutions above, it'll
// default to the YearResolution.
//
// The resolutions are exact and does not take into account leap seconds etc.
func GetUserRoastCountTimeseries(start, end time.Time, resolution time.Duration, username string) (
	timeseries RoastCountTimeseries, err error) {

	return getRoastCountTimeseries(start, end, resolution, username)
}

// getRoastCountTimeseries returns a timeseries of number of Roasts per time
// unit. See: GetGlobalRoastCountTimeseries or GetUserRoastCountTimeseries.
func getRoastCountTimeseries(start, end time.Time, resolution time.Duration, username string) (
	timeseries RoastCountTimeseries, err error) {

	sqlInterval, sqlResolution := getSQLResolution(resolution)
	sqlStart := formatSQLTime(start)
	sqlEnd := formatSQLTime(end)

	rows, err := database.Query(`
		WITH "time_series" AS (
			SELECT generate_series(
				$1::timestamp, -- Start point date with time.
				$2::timestamp, -- End point date with time.
				$3::interval   -- Time series resolution.
			) AS "datapoint"
		)

		SELECT
			"time_series"."datapoint",
			count(distinct r."id")
		FROM "time_series"
		LEFT JOIN "roaster"."roast" AS r
			-- Truncate and compare per resolution.
			ON date_trunc($4, r."create_time") = "time_series"."datapoint"
			AND COALESCE(TRIM($5), '')='' OR LOWER(username)=LOWER(TRIM($5))
		GROUP BY 1
	`, sqlStart, sqlEnd, sqlInterval, sqlResolution, username)
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

// getSQLResolution returns the SQL interval and resolution for the provided
// resolution.
func getSQLResolution(resolution time.Duration) (sqlInterval, sqlResolution string) {
	var resIdentifier string
	switch resolution {
	case SecondResolution:
		resIdentifier = "second"
	case MinuteResolution:
		resIdentifier = "minute"
	case HourResolution:
		resIdentifier = "hour"
	case DayResolution:
		resIdentifier = "day"
	case YearResolution:
	default:
		resIdentifier = "year"
	}

	sqlInterval = fmt.Sprintf("%d %s", 1, resIdentifier)
	sqlResolution = resIdentifier

	return
}

// formatSQLTime returns the time formatted as RFC3339 w/ decimals.
func formatSQLTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000000Z")
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
