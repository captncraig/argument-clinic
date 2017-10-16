package main

import (
	"fmt"

	"github.com/captncraig/argument-clinic/data/sqlite"
)

func main() {
	d, err := sqlite.New("data.db?_foreign_keys=1")
	fmt.Println(d, err)
}
