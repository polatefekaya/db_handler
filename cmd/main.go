package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//var app config.AppConfig

	var connStr string
	connStr = "postgres://postgres:root@localhost:5432/goptest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
