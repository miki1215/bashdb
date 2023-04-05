package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type filename struct {
	filename string
}

func LineCheck(filename string) int {
	openedfile, _ := os.Open("ba")
	defer openedfile.Close()
	scanner := bufio.NewScanner(openedfile)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines
}
func parseFile(filename string, ch chan string) {
	openedfile, _ := os.Open(filename)
	defer openedfile.Close()
	scanner := bufio.NewScanner(openedfile)
	linennumber := 0
	for scanner.Scan() {
		linennumber++
		line := string(strconv.Itoa(linennumber) + " " + scanner.Text())
		ch <- line
	}
	close(ch)
}

func measureTime(f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return elapsed
}

func ReadFromChan(ch chan string) {
	for {
		select {
		case line := <-ch:
			fmt.Println(line)

			time.Sleep(1 * time.Second)
		default:
			fmt.Println(" Not")
			time.Sleep(1 * time.Second)
		}
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
	stmt, err := db.Prepare("INSERT INTO my_table (name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := db.Query("SELECT * FROM my_table")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the received results from the databese to the console
	//	for rows.Next() {
	//		var id int
	//		var name string
	//		var date string
	//		err = rows.Scan(&id, &name, &date)
	//		fmt.Printf("id: %d, name: %s, date:%v", id, name, date)
	//	}
	//
	lines := LineCheck("ba")
	ch := make(chan string, lines)
	go parseFile("ba", ch)
	ReadFromChan(ch)
}
