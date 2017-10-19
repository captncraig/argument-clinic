package main

import (
	"log"

	"github.com/captncraig/argument-clinic/data/sqlite"
	"github.com/captncraig/argument-clinic/web"
)

func main() {
	d, err := sqlite.New("data.db?_foreign_keys=1", "site.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(web.Listen(":8787", d))
}
