package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db_url := "postgres://swordemon:@localhost:5432/uddhanthodu?sslmode=disable"

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")
}
