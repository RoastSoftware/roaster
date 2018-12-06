// Package model roast implementation.
package model

import (
	"database/sql"
	"log"
	"time"
)

// RoastMessage represents a general Roast message.
type RoastMessage struct {
	Hash        []byte `json:"hash"`
	Row         uint   `json:"row"`
	Column      uint   `json:"column"`
	Engine      string `json:"engine"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
func (r *RoastResult) AddError(hash []byte, row, column uint, engine, name, description string) {
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
func (r *RoastResult) AddWarning(hash []byte, row, column uint, engine, name, description string) {
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
			log.Println(tx.Rollback())
			return
		}
		err = tx.Commit()
	}()

	roastInsertResult, err := tx.Exec(`
		INSERT INTO "roast"
		(username, code, score, language, create_time)
		VALUES
		(@username, @code, @score, @language, @create_time)`,
		sql.Named("username", roast.Username),
		sql.Named("code", roast.Code),
		sql.Named("score", roast.Score),
		sql.Named("language", roast.Language),
		sql.Named("create_time", roast.CreateTime))
	if err != nil {
		return
	}

	roastID, err := roastInsertResult.LastInsertId()
	if err != nil {
		return
	}

	errorInsertStmt, err := tx.Prepare(`
		INSERT INTO "error"
		VALUES
		(@hash, @row, @column, @engine, @name, @description)`)
	if err != nil {
		return
	}
	defer errorInsertStmt.Close()

	roastHasErrorsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roast_has_errors"
		VALUES
		(@roast, @error)`)
	if err != nil {
		return
	}
	defer roastHasErrorsInsertStmt.Close()

	for _, errorMessage := range roast.Errors {
		errorInsertResult, err := errorInsertStmt.Exec(
			sql.Named("hash", errorMessage.Hash),
			sql.Named("row", errorMessage.Row),
			sql.Named("column", errorMessage.Column),
			sql.Named("engine", errorMessage.Engine),
			sql.Named("name", errorMessage.Name),
			sql.Named("description", errorMessage.Description))
		if err != nil {
			return err
		}

		errorID, err := errorInsertResult.LastInsertId()
		if err != nil {
			return err
		}

		_, err = roastHasErrorsInsertStmt.Exec(
			sql.Named("roast", roastID),
			sql.Named("error", errorID))
		if err != nil {
			return err
		}
	}

	warningInsertStmt, err := tx.Prepare(`
		INSERT INTO "warning"
		VALUES
		(@hash, @row, @column, @engine, @name, @description)`)
	if err != nil {
		return
	}
	defer warningInsertStmt.Close()

	roastHasWarningsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roast_has_warnings"
		VALUES
		(@roast, @warning)`)
	if err != nil {
		return
	}
	defer roastHasWarningsInsertStmt.Close()

	for _, warningMessage := range roast.Warnings {
		warningInsertResult, err := warningInsertStmt.Exec(
			sql.Named("hash", warningMessage.Hash),
			sql.Named("row", warningMessage.Row),
			sql.Named("column", warningMessage.Column),
			sql.Named("engine", warningMessage.Engine),
			sql.Named("name", warningMessage.Name),
			sql.Named("description", warningMessage.Description))
		if err != nil {
			return err
		}

		warningID, err := warningInsertResult.LastInsertId()
		if err != nil {
			return err
		}

		_, err = roastHasWarningsInsertStmt.Exec(
			sql.Named("roast", roastID),
			sql.Named("warning", warningID))
		if err != nil {
			return err
		}
	}

	return
}
