package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type file struct {
	filename string
	line     string
	db       *dbs
}

type dbs struct {
	db *sql.DB
	wg *sync.WaitGroup
}

type up interface {
	UploadDb()
	UploadRecord()
}

func (f file) parseFile() {
	openedfile, _ := os.Open(f.filename)
	defer openedfile.Close()
	//ch := make(chan string)
	scanner := bufio.NewScanner(openedfile)
	for scanner.Scan() {
		f.line = scanner.Text()
		go f.UploadRecord()
	}
}

func measureTime(f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return elapsed
}

func (f *file) UploadRecord() {
	stmt, err := f.db.db.Prepare("INSERT INTO my_table (name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// execute prepared statement with values
	stmt.Exec(f.line)
	f.db.wg.Done()
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

	// RECEIVING AND READING

	// Execute a SELECT query and retrieve the results
	rows, err := db.Query("SELECT * FROM my_table")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the received results from the databese to the console
	for rows.Next() {
		var id int
		var name string
		var date string
		err = rows.Scan(&id, &name, &date)
		fmt.Printf("id: %d, name: %s, date:%v", id, name, date)
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	file := file{filename: "ba", db: &dbs{db: db, wg: &wg}}
	file.parseFile()
	wg.Wait()

}
