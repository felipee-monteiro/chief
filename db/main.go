package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

type DatabaseConnection struct {
	instance *sql.DB
	dsn      *url.URL
	rawDsn   string
}

func (d *DatabaseConnection) Connect(username, password, server, database string, port int) bool {

	if d.instance != nil {
		return true
	}

	query := url.Values{}
	query.Add("database", database)
	query.Add("encrypt", "disable")
	query.Add("TrustServerCertificate", "true")

	rawQ := query.Encode()

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%d", server, port),
		RawQuery: rawQ,
	}

	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		log.Fatal(err)
	}

	d.dsn = u
	d.rawDsn = rawQ

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	d.instance = db

	return false
}

func (d *DatabaseConnection) GetInstance() (*sql.DB, bool) {
	var dbInstance = d.instance

	if dbInstance == nil {
		return nil, false
	}

	return dbInstance, true
}

func (d *DatabaseConnection) Query(q string) *sql.Rows {
	dbInstance, ok := d.GetInstance()

	if !ok {
		panic("Database not connected, please try again")
	}

	result, err := dbInstance.Query(q)

	if err != nil {
		log.Fatal("Erro na execução da query: ", err)
	}

	return result
}

func (d *DatabaseConnection) GetRawDsn() string {
	return d.rawDsn
}
