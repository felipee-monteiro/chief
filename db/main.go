package db

import (
	"database/sql"
	"fmt"
	"log"
    	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)  

func Connect(username, password, server, database string, port int) *sql.DB {
    db_username := username
    db_password := password
    db_database := database
    db_server   := server
    db_port     := port

    query := url.Values{}
    query.Add("database", db_database)
    query.Add("encrypt", "disable")
    query.Add("TrustServerCertificate", "true") 

    u := &url.URL{
	Scheme:   "sqlserver",
	User:     url.UserPassword(db_username, db_password),
	Host:     fmt.Sprintf("%s:%d", db_server, db_port),
	RawQuery: query.Encode(),
    }

    db, err := sql.Open("sqlserver", u.String())
	
    if err != nil {
	log.Fatal(err)
    }

    err = db.Ping()
	
    if err != nil {
	log.Fatal(err)
    }

    return db
}

func Query(q string, conn *sql.DB) *sql.Rows {		
	result, err := conn.Query(q)

	if err != nil {
		log.Fatal("Erro na execução da query: ", err);
	}

	return result
}
