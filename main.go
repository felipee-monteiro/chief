package main

import (
	"log"
	"net/http"

	"api-test/routes"
	"api-test/db"
)

func main() {	
	conn := db.Connect()
	routes.InitialPage(conn)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
