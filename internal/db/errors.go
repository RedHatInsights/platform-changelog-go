package db

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

// handling sql errors here instead of the web layer
func evaluateError(err error) error {
	if err == nil {
		return nil
	}

	// https://gorm.io/docs/error_handling.html#ErrRecordNotFound
	if err == gorm.ErrRecordNotFound || err == sql.ErrNoRows {
		return ErrNotFound
	}

	return err
}
