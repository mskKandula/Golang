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

func updateRow(id, inc int, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := db.Exec("UPDATE Test SET val = val+ ? WHERE id = ?", inc, id)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	// Create the Test table if it doesn't exist
	// _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Test (
	// 	id INT PRIMARY KEY,
	// 	val INT
	// )`)
	// if err != nil {
	// 	panic(err)
	// }

	// Insert initial data into the table
	// _, err := db.Exec("INSERT INTO Test (id, val) VALUES (?, ?)", 1, 0)
	// if err != nil {
	// 	panic(err)
	// }

	// Set the isolation level for the session
	_, err := db.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE")
	if err != nil {
		panic(err)
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Number of concurrent goroutines
	numGoroutines := 10

	// Launch concurrent goroutines to update the row
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go updateRow(2, 10, &wg) // Increment the value by 1 in each update
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Fetch and print the final value after updates
	var finalVal int
	err = db.QueryRow("SELECT val FROM Test WHERE id = ?", 2).Scan(&finalVal)
	if err != nil {
		panic(err)
	}
	fmt.Println("Final value:", finalVal)
}
