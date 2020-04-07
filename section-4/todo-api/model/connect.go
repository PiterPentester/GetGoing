package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "username:pass@(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Connected to DB")
	con = db
	return db
}
