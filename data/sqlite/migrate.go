package sqlite

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/captncraig/argument-clinic/data"
	"github.com/rubenv/sql-migrate"
)

func (d *db) migrate() error {
	d.Lock()
	defer d.Unlock()
	migs := []*migrate.Migration{}
	for i, mig := range []string{
		mig0CreateSite,
		mig1CreateHosts,
		mig2InsertDefaultSite,
		mig3InsertDefaultHostName,
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

const mig0CreateSite = `CREATE TABLE sites ( 
	SiteID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
	Salt BLOB NOT NULL UNIQUE 
	);`

const mig1CreateHosts = `CREATE TABLE hostnames ( 
	Name TEXT NOT NULL UNIQUE, 
	SiteID INTEGER NOT NULL, 
	PRIMARY KEY(Name), 
	FOREIGN KEY(SiteID) REFERENCES sites(SiteID) )`

var mig2InsertDefaultSite = func() string {
	return fmt.Sprintf(`INSERT INTO sites(Salt) VALUES(X'%s')`, hex.EncodeToString(data.GenerateSalt()))
}()

const mig3InsertDefaultHostName = `INSERT INTO hostnames(SiteID, Name) VALUES(1, '*')`
