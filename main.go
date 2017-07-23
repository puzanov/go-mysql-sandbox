package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	// check for driver
	if err != nil {
		log.Fatal(err)
	}

	//check for connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}