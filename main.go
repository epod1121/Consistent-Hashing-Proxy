package main

import (
	"database/sql"
	"fmt"
	"log"
	//_ "github.com/lib/pq"
)

func main() {

	connString := "host=localhost port=5000 user=postgres password=dissys dbname=hashdb sslmode = disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed: ", err)
	}

	fmt.Println("Connected to Database!")
}