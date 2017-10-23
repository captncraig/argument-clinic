package sqlite

import (
	"fmt"
	"log"

	"github.com/rubenv/sql-migrate"
)

func (d *db) migrate() error {
	migs := []*migrate.Migration{}
	for i, mig := range []string{
		mig0CreateComments,
	} {
		migs = append(migs, &migrate.Migration{
			Id: fmt.Sprint(i),
			Up: []string{mig},
		})
	}
	migrations := &migrate.MemoryMigrationSource{
		Migrations: migs,
	}

	n, err := migrate.Exec(d.db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}
	if n > 0 {
		log.Printf("Applied %d migrations!\n", n)
	}
	return nil
}

const mig0CreateComments = `
CREATE TABLE Comments (
	ID integer PRIMARY KEY AUTOINCREMENT,
	Name text,
	Email text,
	Text text NOT NULL,
	Created integer(4) NOT NULL DEFAULT (strftime('%s','now'))
);
`
