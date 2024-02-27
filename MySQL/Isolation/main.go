package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername = "root"
	dbPassword = "root"
	dbName     = "oes"
	dbHost     = "localhost:3306"
)

var (
	db *sql.DB
)

func initDB() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbHost, dbName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func updateRow(id, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		query *sql.Stmt
		err   error
	)
	if i%2 == 0 {
		query, err = db.Prepare("UPDATE Test SET val = val + 100 WHERE id=?")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		query, err = db.Prepare("UPDATE Test SET val = val + 50 WHERE id=?")
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err = query.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Number of concurrent goroutines
	numGoroutines := 100

	// Launch concurrent goroutines to update the row
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go updateRow(2, i, &wg) // Update row with ID = 1
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All updates completed.")
}
