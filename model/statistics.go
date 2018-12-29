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

// RoastRatio represents a ratio between lines of code versus errors and
// warnings.
type RoastRatio struct {
	LinesOfCode      uint64 `json:"linesOfCode"`
	NumberOfErrors   uint64 `json:"numberOfErrors"`
	NumberOfWarnings uint64 `json:"numberOfWarnings"`
}

// GetRoastTimeseries returns a timeseries of number of Roasts per time
// unit. Username is optional, an empty string represent every user.
//
// The minimum interval is 1 minute, anything less will be set to 1 minute per
// default.
func GetRoastTimeseries(start, end time.Time, interval time.Duration, username string) (
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

// GetLinesOfCode returns the number of lines of code for everyone or a specific
// user. An empty string as username represents everyone.
func GetLinesOfCode(username string) (lines LinesOfCode, err error) {
	err = database.QueryRow(`
		SELECT COALESCE(SUM("lines_of_code"), 0) AS "lines_of_code"
		FROM roaster.roast_statistics AS s
		JOIN roaster.roast AS r
		ON r.id = s.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(r.username)=LOWER(TRIM($1))
	`, username).Scan(&lines.Lines)
	return
}

// GetNumberOfRoasts returns the number of Roasts for everyone or a specific
// user. An empty string as username represents everyone.
func GetNumberOfRoasts(username string) (numberOfRoasts NumberOfRoasts, err error) {
	err = database.QueryRow(`
		SELECT COUNT(username)
		FROM roaster.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(username)=LOWER(TRIM($1))
	`, username).Scan(&numberOfRoasts.Count)
	return
}

// GetRoastRatio returns the lines of code, number of errors and warnings for
// everyone or a specific user. An Empty string as usernmae represents everyone.
func GetRoastRatio(username string) (roastRatio RoastRatio, err error) {
	err = database.QueryRow(`
		SELECT
			COALESCE(SUM(s."number_of_errors"), 0) AS "number_of_errors",
			COALESCE(SUM(s."number_of_warnings"), 0) AS "number_of_warnings",
			COALESCE(SUM("lines_of_code"), 0) AS "lines_of_code"
		FROM roaster.roast_statistics AS s
			JOIN roaster.roast AS r
			ON r.id = s.roast
		WHERE COALESCE(TRIM($1), '')='' OR
		      LOWER(r.username)=LOWER(TRIM($1))
	`, username).Scan(&roastRatio.NumberOfErrors, &roastRatio.NumberOfWarnings, &roastRatio.LinesOfCode)
	return
}
