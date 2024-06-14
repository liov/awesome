package main

import (
	"github.com/hopeio/cherry/initialize/conf_dao/duckdb"
	"log"
)

func main() {
	config := duckdb.Config{
		DSN:         "./duck.db?access_mode=read_only&threads=4",
		Path:        "",
		AccessMode:  "",
		Threads:     0,
		BootQueries: nil,
	}
	db, err := config.Build()
	if err != nil {
		log.Fatal("Build err", err)
	}
	_, err = db.Exec(`CREATE TABLE people (id INTEGER, name VARCHAR)`)
	if err != nil {
		log.Fatal("Exec err", err)
	}
}
