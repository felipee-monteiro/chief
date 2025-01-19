package main

import (
	"log"
	"net/http"

	"api-test/routes"
	"api-test/db"
)

func main() {	
	conn := db.Connect("sa", "Epilefac57#$!$24042002", "host.docker.internal", "sigma", 1433)
	routes.InitialPage(conn)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
