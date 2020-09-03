package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b, err := os.Open("C:/Users/Mehdi/Desktop/test.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(b)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	b.Close()

	for _, eachline := range txtlines {
		if strings.HasPrefix(eachline, "#7") {
			t := strings.Replace(eachline, "#", "", -1)
			fmt.Println(t)
			stmt, err := db.Prepare("INSERT INTO test(tel) VALUES(?)")
			if err != nil {
				log.Fatal(err)
			}
			res, err := stmt.Exec(t)
			if err != nil {
				log.Fatal(err)
			}
			lastID, err := res.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}
			rowCnt, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("lastId = %d, affected = %d\n", lastID, rowCnt)

		}
	}

}
