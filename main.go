package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func parseFile(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	connStr := "postgresql://postgres:miki@localhost:5432/bash?sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//	_, err = db.Exec("ALTER TABLE my_table ADD COLUMN created_at TIMESTAMP DEFAULT NOW()")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	// use prepared statement to insert data into table
	stmt, err := db.Prepare("INSERT INTO my_table (name, id) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// execute prepared statement with values
	stmt.Exec("jjjj", 20000)

	// Execute a SELECT query and retrieve the results
	rows, err := db.Query("SELECT * FROM my_table")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Print the results to the console
	for rows.Next() {
		var id int
		var name string
		var date string
		err = rows.Scan(&id, &name, &date)
		fmt.Printf("id: %d, name: %s\n, date:%v\n", id, name, date)
	}

	parseFile("ba")

}
