// Package model roast implementation.
package model

import (
	"log"
	"time"

	"github.com/satori/go.uuid"
)

// RoastMessage represents a general Roast message.
type RoastMessage struct {
	Hash        uuid.UUID `json:"hash"`
	Row         uint      `json:"row"`
	Column      uint      `json:"column"`
	Engine      string    `json:"engine"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// RoastError represents a Roast error message.
type RoastError struct {
	RoastMessage
}

// RoastWarning represents a Roast error message.
type RoastWarning struct {
	RoastMessage
}

// RoastResult represent a Roast result.
type RoastResult struct {
	Username   string         `json:"username"`
	Code       string         `json:"code"`
	Score      uint           `json:"score"`
	Language   string         `json:"language"`
	Errors     []RoastError   `json:"errors"`
	Warnings   []RoastWarning `json:"warnings"`
	CreateTime time.Time      `json:"create_time"`
}

// AddError adds an error to the RoastResult.
func (r *RoastResult) AddError(hash uuid.UUID, row, column uint, engine, name, description string) {
	r.Errors = append(r.Errors, RoastError{
		RoastMessage{
			Hash:        hash,
			Row:         row,
			Column:      column,
			Engine:      engine,
			Name:        name,
			Description: description,
		},
	})
}

// AddWarning adds an warning to the RoastResult.
func (r *RoastResult) AddWarning(hash uuid.UUID, row, column uint, engine, name, description string) {
	r.Warnings = append(r.Warnings, RoastWarning{
		RoastMessage{
			Hash:        hash,
			Row:         row,
			Column:      column,
			Engine:      engine,
			Name:        name,
			Description: description,
		},
	})
}

// PutRoast adds a RoastResult to the database.
func PutRoast(roast RoastResult) (err error) {
	tx, err := database.Begin()
	if err != nil {
		return
	}
	defer func() { // Rollback transaction on error.
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Println(rollBackErr)
			}

			return
		}
		err = tx.Commit()
	}()

	var roastID int
	err = tx.QueryRow(`
		INSERT INTO "roast"
		(username, code, score, language, create_time)
		VALUES
		(TRIM($1), $2, $3, $4, $5)
 		RETURNING id
		`,
		roast.Username,
		roast.Code,
		roast.Score,
		roast.Language,
		time.Now()).Scan(&roastID)
	if err != nil {
		return
	}

	errorInsertStmt, err := tx.Prepare(`
		INSERT INTO "error"
		(hash, row, "column", engine, name, description)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id`)
	if err != nil {
		return
	}
	defer errorInsertStmt.Close()

	roastHasErrorsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roast_has_errors"
		(roast, error)
		VALUES
		($1, $2)`)
	if err != nil {
		return
	}
	defer roastHasErrorsInsertStmt.Close()

	for _, errorMessage := range roast.Errors {
		var errorID int

		err := errorInsertStmt.QueryRow(
			errorMessage.Hash,
			errorMessage.Row,
			errorMessage.Column,
			errorMessage.Engine,
			errorMessage.Name,
			errorMessage.Description).Scan(&errorID)
		if err != nil {
			return err
		}

		_, err = roastHasErrorsInsertStmt.Exec(roastID, errorID)
		if err != nil {
			return err
		}
	}

	warningInsertStmt, err := tx.Prepare(`
		INSERT INTO "warning"
		(hash, row, "column", engine, name, description)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id`)
	if err != nil {
		return
	}
	defer warningInsertStmt.Close()

	roastHasWarningsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roast_has_warnings"
		(roast, warning)
		VALUES
		($1, $2)`)
	if err != nil {
		return
	}
	defer roastHasWarningsInsertStmt.Close()

	for _, warningMessage := range roast.Warnings {
		var warningID int

		err := warningInsertStmt.QueryRow(
			warningMessage.Hash,
			warningMessage.Row,
			warningMessage.Column,
			warningMessage.Engine,
			warningMessage.Name,
			warningMessage.Description).Scan(&warningID)
		if err != nil {
			return err
		}

		_, err = roastHasWarningsInsertStmt.Exec(roastID, warningID)
		if err != nil {
			return err
		}
	}

	return
}
