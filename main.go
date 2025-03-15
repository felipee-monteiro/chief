package main

import (
	"chief/db"
)

func main() {
	db := db.DatabaseConnection{}
	db.Connect("sa", "Epilefac57#$!$24042002", "host.docker.internal", "sigma", 1433)
}
