package sqlite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/captncraig/argument-clinic/data"
	"github.com/captncraig/argument-clinic/models"
)

type db struct {
	db *sql.DB
}

// New creates a new sql data store, and ensures that all data migrations are run.
func New(sqlFile string) (data.DataAccess, error) {
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

func (d *db) CreateComment(ctx context.Context, req *models.CreateCommentRequest) (uint64, error) {
	const sql = `
	INSERT INTO Comments(Name, Email, Text)
	VALUES (?,?,?);`
	r, err := d.db.ExecContext(ctx, sql, req.Name, req.Email, req.Text)
	if err != nil {
		return 0, err
	}
	last, err := r.LastInsertId()
	return uint64(last), err
}
