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
func GetRoastTimeseries(start, end time.Time, interval time.Duration, username string, followees bool) (
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
			t."datapoint" AS "timestamp",
			COUNT(r."id") AS "count",
			COALESCE(SUM(s."number_of_errors"), 0) AS "number_of_errors",
			COALESCE(SUM(s."number_of_warnings"), 0) AS "number_of_warnings",
			COALESCE(SUM(s."lines_of_code"), 0) AS "lines_of_code"
		FROM "time_series" AS t

		-- Collect users followees if requested.
		LEFT OUTER JOIN "roaster"."user_followees" AS f
			ON $6 AND LOWER(f."username")=LOWER(TRIM($5))

		-- Collect all the Roasts per resolution.
		LEFT JOIN "roaster"."roast" AS r
			-- Truncate and compare per resolution.
			ON
				date_trunc('minute', "roaster".round_minutes(r."create_time"::timestamp, $4)) = t."datapoint"
				AND (COALESCE(TRIM($5), '')='' OR
				NOT $6 AND LOWER(r."username")=LOWER(TRIM($5)) OR -- Optionally only return for specified user.
				$6 AND r."username" = f."followee") -- Or only return that users followees.

			-- Collect statistics for the Roasts.
			LEFT JOIN "roaster"."roast_statistics" AS s
				ON s."roast" = r."id"

		GROUP BY 1
		ORDER BY "timestamp" DESC -- First in result is the latest timestamp.
		LIMIT 1000 -- Do not return more than 1000 rows back in time.
	`,
		start,
		end,
		fmt.Sprintf("%.0f minutes", interval.Minutes()),
		interval.Minutes(),
		username,
		followees)
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

// GetLinesOfCode returns the number of lines of code for everyone, a specific
// user or that user's followees. An empty string as username represents everyone.
func GetLinesOfCode(username string, followees bool) (lines LinesOfCode, err error) {
	err = database.QueryRow(`
		SELECT COALESCE(SUM("lines_of_code"), 0) AS "lines_of_code"
		FROM "roaster"."roast_statistics" AS s

		-- Optionally collect the users followees.
		LEFT OUTER JOIN "roaster"."user_followees" AS f
			ON $2 AND LOWER(f."username")=LOWER(TRIM($1))

		-- Collect Roasts to compare with specfic username.
		JOIN "roaster"."roast" AS r
			ON r."id" = s."roast"

		-- Return either globally, for specific user, or for specific users followees.
		WHERE COALESCE(TRIM($1), '')='' OR
		      NOT $2 AND LOWER(r."username")=LOWER(TRIM($1)) OR
		      $2 AND r."username" = f."followee"
	`, username, followees).Scan(&lines.Lines)
	return
}

// GetNumberOfRoasts returns the number of Roasts for everyone, a specific
// user, or that user's followees. An empty string as username represents everyone.
func GetNumberOfRoasts(username string, followees bool) (numberOfRoasts NumberOfRoasts, err error) {
	err = database.QueryRow(`
		SELECT COUNT(r."username")
		FROM roaster.roast AS r

		-- Optionally collect the users followees.
		LEFT OUTER JOIN "roaster"."user_followees" AS f
			ON $2 AND LOWER(f."username")=LOWER(TRIM($1))

		-- Return either globally, for specific user, or for specific users followees.
		WHERE COALESCE(TRIM($1), '')='' OR
		      (NOT $2 AND LOWER(r."username")=LOWER(TRIM($1))) OR
		      ($2 AND r."username" = f."followee")
	`, username, followees).Scan(&numberOfRoasts.Count)
	return
}

// GetRoastRatio returns the lines of code, number of errors and warnings for a
// specific user.
func GetRoastRatio(username string) (roastRatio RoastRatio, err error) {
	err = database.QueryRow(`
		SELECT
			COALESCE(SUM(s."number_of_errors"), 0) AS "number_of_errors",
			COALESCE(SUM(s."number_of_warnings"), 0) AS "number_of_warnings",
			COALESCE(SUM("lines_of_code"), 0) AS "lines_of_code"
		FROM roaster.roast_statistics AS s
			JOIN roaster.roast AS r
			ON r.id = s.roast
		WHERE LOWER(r.username)=LOWER(TRIM($1))
	`, username).Scan(&roastRatio.NumberOfErrors, &roastRatio.NumberOfWarnings, &roastRatio.LinesOfCode)
	return
}
