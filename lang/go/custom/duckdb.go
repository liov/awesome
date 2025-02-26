package main

import (
	"github.com/hopeio/initialize/dao/duckdb"
	"log"
	"time"
)

// CGO_CFLAGS=-ID:/sdk/libduckdb-windows-amd64;CGO_ENABLED=1;CGO_LDFLAGS=-LD:/sdk/libduckdb-windows-amd64 -tags=duckdb_use_lib,go1.22
type Log struct {
	Time    time.Time
	Level   string
	TraceId int64
	Caller  string
	Message string
}

func main() {
	config := duckdb.Config{
		DSN:         "./duck.db?access_mode=READ_WRITE&threads=4",
		Path:        "",
		AccessMode:  "",
		Threads:     0,
		BootQueries: nil,
	}
	db, err := config.Build()
	if err != nil {
		log.Fatal("Build err", err)
	}
	/*	_, err = db.Exec(`CREATE TYPE level AS ENUM ('debug', 'info', 'warn','error')`)
		if err != nil {
			log.Fatal("CREATE TYPE err", err)
		}*/
	/*	_, err = db.Exec(`DROP TABLE logs`)
		if err != nil {
			log.Println("DROP TABLE err", err)
		}
		_, err = db.Exec(`CREATE TABLE logs (traceId UBIGINT, time Timestamp,level level, caller VARCHAR, message TEXT)`)
		if err != nil {
			log.Fatal("CREATE TABLE err", err)
		}
		_, err = db.Exec(`INSERT INTO logs SELECT * FROM read_json_auto('D:/work/agent-1717641277076.log', format = 'newline_delimited',ignore_errors = true, columns = {traceId: 'UBIGINT', time: 'Timestamp',level :'level', caller: 'VARCHAR', message: 'TEXT'})`)
		if err != nil {
			log.Fatal("INSERT err", err)
		}*/
	rows, err := db.Query(`SELECT * FROM logs LIMIT 5;`)
	if err != nil {
		log.Fatal("SELECT err", err)
	}
	for rows.Next() {
		var log1 Log
		err = rows.Scan(&log1.TraceId, &log1.Time, &log1.Level, &log1.Caller, &log1.Message)
		if err != nil {
			log.Fatal("Scan err", err)
		}
		log.Println(log1)
	}
}
