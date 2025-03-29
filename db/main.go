package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBCache struct {
	connection *sql.DB
}

var cache = DBCache{}

func Connect() *sql.DB {

	if cache.connection != nil {
		return cache.connection
	}

	// TODO: transform into absolute path
	db, err := sql.Open("sqlite3", "../chief.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	cache.connection = db

	return db
}
