package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"database/sql"
	"fmt"
	"log"
    "net/url"
    "strings"

	_ "github.com/microsoft/go-mssqldb"
)  

//export Connect
func Connect(username, password, server, database string, port int) *C.char {
	
    valuesToValidate := [4]string{username, password, server, database}

    for _, v := range valuesToValidate {
        if len(strings.TrimSpace(v)) == 0 {                                                                      log.Fatal("Argumentos insuficientes.")                                                         }          
    }

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
		return C.CString("Erro ao abrir a conexão com o banco de dados")
	}

	err = db.Ping()
	
    if err != nil {
		log.Fatal(err)
		return C.CString("Erro ao verificar a conexão com o banco de dados")
	}

	return C.CString("Conexão com o banco de dados estabelecida com sucesso")
}

func Query(q string, conn *sql.DB) *sql.Rows {		
	result, err := conn.Query(q)

	if err != nil {
		log.Fatal("Erro na execução da query: ", err);
	}

	return result
}

func main() {}
