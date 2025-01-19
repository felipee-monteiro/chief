package db

import (
    "C"

	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)  


func Connect(username, password, server, database string, port int) *sql.DB {
	 query := url.Values{}
	 query.Add("database", database)

	 u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%d", server, port),
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

	 log.Println("Conexão com o banco de dados estabelecida")

	 return db
}

func Query(q string, conn *sql.DB) *sql.Rows {		
	result, err := conn.Query(q)

	if err != nil {
		log.Fatal("Erro na execução da query: ", err);
	}

	return result
}
