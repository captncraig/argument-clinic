package sqlite

import (
	"fmt"
	"log"

	"github.com/rubenv/sql-migrate"
)

func (d *db) migrate() error {
	d.Lock()
	defer d.Unlock()
	migs := []*migrate.Migration{}
	for i, mig := range []string{
		mig0CreateSite,
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

const mig0CreateSite = `CREATE TABLE 'site' ( 
	'SiteID' INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
	'Salt' BLOB NOT NULL UNIQUE 
	);`
