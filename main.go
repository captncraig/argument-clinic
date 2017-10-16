package main

import (
	"fmt"

	"github.com/captncraig/arguments/data/sqlite"
)

func main() {
	d, err := sqlite.New("data.db")
	fmt.Println(d, err)
}
