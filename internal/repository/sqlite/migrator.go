package sqlite

import (
	"database/sql"
	"github.com/pressly/goose"
)

const dialect = "sqlite3"

func Migrate(db *sql.DB) (err error) {
	if err = goose.SetDialect(dialect); err != nil {
		return
	}

	err = goose.Up(db, ".")
	if err != nil {
		return
	}

	return
}