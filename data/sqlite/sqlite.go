package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/captncraig/argument-clinic/data"
)

type db struct {
	db *sql.DB
}

// New creates a new sql data store, and ensures that all data migrations are run.
func New(sqlFile string, siteConfigFile string) (data.DataAccess, error) {
	d := &db{}
	var err error
	d.db, err = sql.Open("sqlite3", sqlFile)
	if err != nil {
		return nil, err
	}
	if err = d.migrate(); err != nil {
		return nil, err
	}
	return d, nil
}
