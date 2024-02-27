package main

import (
	"database/sql"
	"fmt"
	"time"

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

// ... (previous code)

func updateRowInTransaction(tx *sql.Tx, id, val int) {
	// defer wg.Done()

	_, err := tx.Exec("UPDATE Test SET val = val+? WHERE id = ?", val, id)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	t1 := time.Now()
	initDB()
	defer db.Close()

	// Create a wait group to synchronize goroutines
	// var wg sync.WaitGroup

	// Number of concurrent goroutines
	numGoroutines := 100

	// Set the isolation level for the session
	_, err := db.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")
	if err != nil {
		panic(err)
	}

	// Begin a transaction with the chosen isolation level
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // Rollback if not committed explicitly

	// Launch concurrent goroutines within the transaction to update the row
	for i := 0; i < numGoroutines; i++ {
		// wg.Add(1)
		updateRowInTransaction(tx, 1, 1) // Increment the value by 1 in each update
	}

	// Wait for all goroutines to finish
	// wg.Wait()

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println("Commit error:", err)
	}

	// Fetch and print the final value after updates
	var finalVal int
	err = db.QueryRow("SELECT val FROM Test WHERE id = ?", 1).Scan(&finalVal)
	if err != nil {
		panic(err)
	}
	fmt.Println("Final value:", finalVal)
	fmt.Println(time.Since(t1))
}
