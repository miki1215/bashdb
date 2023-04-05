package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
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
	openedfile, err := os.Open(filename)
	if err != nil {
		log.Panic("Non-existent file")
	}
	defer openedfile.Close()
	scanner := bufio.NewScanner(openedfile)
	linennumber := 0
	for scanner.Scan() {
		linennumber++
		line := string(scanner.Text())
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
elso:
	for {

		select {
		case line, ok := <-ch:
			if !ok {
				fmt.Println("channel closed, breaking...")
				break elso
			}
			//	fmt.Println(line)
			Upload(line)
		default:
			fmt.Println(" Not")
		}
	}
}
func GetFileName() []string {
	filename := os.Args

	if len(filename) > 1 {
		arg := filename[1]
		fmt.Println("Filename:", arg)
	} else {
		fmt.Println("No argument provided")
	}
	return filename
}
func Upload(name string) {
	connStr := "postgresql://postgres:miki@localhost:5432/bash?sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO my_table (name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(name)
	defer stmt.Close()

}
func main() {
	//	_, err = db.Exec("ALTER TABLE my_table ADD COLUMN created_at TIMESTAMP DEFAULT NOW()")
	// Print the received results from the databese to the console
	//	for rows.Next() {
	//		var id int
	//		var name string
	//		var date string
	//		err = rows.Scan(&id, &name, &date)
	//		fmt.Printf("id: %d, name: %s, date:%v", id, name, date)
	//	}
	//
	filename := GetFileName()
	lines := LineCheck(filename[1])
	ch := make(chan string, lines)
	parseFile(filename[1], ch)
	ReadFromChan(ch)
}
