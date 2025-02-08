package main

import (
	"database/sql"
	"fmt"
	"log"
    	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)  


func Connect(username, password, server, database *C.char, port C.int) *C.char {
    db_username := C.GoString(username)
    db_password := C.GoString(password)
    db_database := C.GoString(database)
    db_server   := C.GoString(server)
    db_port     := int(port)

    query := url.Values{}
	query.Add("database", db_database)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(db_username, db_password),
		Host:     fmt.Sprintf("%s:%d", db_server, db_port),
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
