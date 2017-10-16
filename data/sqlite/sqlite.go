package sqlite

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"github.com/captncraig/arguments/data"
	"github.com/captncraig/arguments/models"
)

type db struct {
	sync.RWMutex
	db *sql.DB
}

// New creates a new sql data store, and ensures that all data migrations are run.
func New(filename string) (data.DataAccess, error) {
	d := &db{}
	var err error
	d.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	if err = d.migrate(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *db) CreateComment(*models.Comment) (uint64, error) {
	return 0, nil
}
