package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgresql://postgres:miki@localhost:5432/bash?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Execute a SELECT query and retrieve the results
	rows, err := db.Query("SELECT * FROM my_table")
	defer rows.Close()

	// Print the results to the console
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}
