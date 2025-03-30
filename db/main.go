package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Connect opens a connection to the "chief.sqlite3" database in the current
// working directory. If the file does not exist, it will be created. If the
// connection has already been opened, it will return the existing connection.
// If any error occurs, it will log the error and exit the program.
func Connect() *sql.DB {
	abs, patherr := filepath.Abs(".")

	if patherr != nil {
		fmt.Println(patherr)
		os.Exit(1)
		return nil
	}

	url := url.URL{
		Scheme: "file",
		Path:   path.Join(abs, "chief.sqlite3"),
		RawQuery: "_pragma=foreign_keys(1)&" +
			"_time_format=sqlite",
	}

	db, err := sql.Open("sqlite3", url.String())

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}

// CreateTables creates the "migrations" table in the database if it does not
// already exist. The table contains columns for an auto-incrementing ID,
// migration name, SQL scripts for applying and rolling back the migration, and
// a timestamp for when the migration was created. If an error occurs during
// table creation, it logs the error and exits the program.
func CreateTables(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			up_sql TEXT NOT NULL,
			down_sql TEXT NOT NULL,
			executed BOOLEAN NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Migrate creates the "migrations" table in the database if it does not already
// exist, as well as any other required tables. It should be called after the
// database connection is established.
func Migrate(db *sql.DB) {
	CreateTables(db)
}

// Close closes the given database connection. If the connection is already
// closed, it does nothing.
func Close(db *sql.DB) {
	db.Close()
}

// MigrateAndClose creates the "migrations" table in the database if it does not
// already exist, as well as any other required tables, and then closes the
// database connection. It should be called after the database connection is
// established.
func MigrateAndClose(db *sql.DB) {
	Migrate(db)
	Close(db)
}

// DropTables drops the "migrations" table in the database if it exists. It
// should generally not be called directly; instead, use DropAndClose to drop
// the tables and then close the database connection. If an error occurs while
// dropping the table, it logs the error and exits the program.
func DropTables(db *sql.DB) {
	sql := `
		DROP TABLE IF EXISTS migrations;
	`

	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// DropAndClose drops the "migrations" table in the database if it exists, and
// then closes the database connection. It should generally not be called
// directly; instead, use DropAndClose to drop the tables and then close the
// database connection. If an error occurs while dropping the table, it logs
// the error and exits the program.
func DropAndClose(db *sql.DB) {
	DropTables(db)
	Close(db)
}

// MigrateAndDrop drops the "migrations" table in the database if it exists,
// and then recreates it. It should generally not be called directly; instead,
// use MigrateAndDrop to drop and recreate the tables. If an error occurs while
// dropping or recreating the table, it logs the error and exits the program.
func MigrateAndDrop(db *sql.DB) {
	DropTables(db)
	Migrate(db)
}

// CreateMigration inserts a new migration record into the "migrations" table.
// The given name, up SQL script, and down SQL script are inserted into the
// table. If any error occurs during insertion, it logs the error and exits the
// program.
func CreateMigration(db *sql.DB, name, up, down string) {
	sql := `
		INSERT INTO migrations (name, up_sql, down_sql, executed)
		VALUES (?, ?, ?, 1);
	`

	if _, err := db.Exec(sql, name, up, down); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// IsExecuted returns true if the migration with the given name has been
// executed, and false otherwise. If an error occurs while querying the
// database, it logs the error and exits the program.
func IsExecuted(db *sql.DB, name string) bool {
	sql := `
		SELECT executed
		FROM migrations
		WHERE name = ?;
	`

	var executed bool
	if err := db.QueryRow(sql, name).Scan(&executed); err != nil {
		return false
	}

	return executed
}
