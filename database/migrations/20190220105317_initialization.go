package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20190220105317, Down20190220105317)
}

func Up20190220105317(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func Down20190220105317(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}