package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type DBCache struct {
	connection *sql.DB
}

var cache = DBCache{}

// Connect opens a connection to the "chief.sqlite3" database in the current
// working directory. If the file does not exist, it will be created. If the
// connection has already been opened, it will return the existing connection.
// If any error occurs, it will log the error and exit the program.
func Connect() *sql.DB {

	if cache.connection != nil {
		return cache.connection
	}

	abs, patherr := filepath.Abs(path.Dir(os.Args[0]))

	if patherr != nil {
		fmt.Println(patherr)
		os.Exit(1)
		return nil
	}

	db, err := sql.Open("sqlite3", path.Join(abs, "chief.sqlite3"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cache.connection = db

	return db
}
