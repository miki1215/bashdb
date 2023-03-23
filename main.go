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

func parseFile(wg *sync.WaitGroup, filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	//ch := make(chan string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		go UploadDb(wg, line)
	}
}

func measureTime(f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return elapsed
}

func UploadDb(wg *sync.WaitGroup, line string) {
	fmt.Println("\n------")
	fmt.Println(line)
	wg.Done()
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
	parseFile(&wg, "ba")
	wg.Wait()

}
