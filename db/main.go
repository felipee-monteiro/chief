package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)  

func Connect() *sql.DB {
	 query := url.Values{}
	 query.Add("database", "sigma")

	 u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "Epilefac57#$!$24042002"),
		Host:     fmt.Sprintf("%s:%d", "host.docker.internal", 1433),
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