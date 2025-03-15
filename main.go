package main

import (
	"fmt"
	"log"
	"net/http"

	"api-test/db"
	"api-test/routes"
)

func main() {
	db := db.DatabaseConnection{}
	db.Connect("sa", "Epilefac57#$!$24042002", "host.docker.internal", "sigma", 1433)

	fmt.Println(db.GetRawDsn())

	routes.InitialPage(&db)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
