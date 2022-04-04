package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Row struct {
	Business   string    `csv:"businessname" db:"business_name"`
	Licstatus  string    `csv:"licstatus" db:"license_status"`
	Result     string    `csv:"result" db:"result"`
	Violdesc   string    `csv:"violdesc" db:"description"`
	Violdttm   time.Time `csv:"violdttm" db:"time"`
	Violstatus string    `csv:"violstatus" db:"status"`
	Viollevel  string    `csv:"viollevel" db:"-"`
	Level      int       `db:"level"`
	Comments   string    `csv:"comments" db:"comments"`
	Address    string    `csv:"address" db:"address"`
	City       string    `csv:"city" db:"city"`
	Zip        string    `csv:"zip" db:"zip"`
}

var (
	db          *sql.DB
	err         error
	insertQuery string
)

func main() {
	file, err := os.Open("boston.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db, err = sql.Open("mysql", "UserName:Password@tcp(IP:Port)/DB")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	insertQuery = "INSERT INTO violations VALUES(?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	_, _, err = ETL(file, stmt)

	duration := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Time Taken to Process: ", duration)
}

func ETL(csvFile io.Reader, stmt *sql.Stmt) (int, int, error) {

}

func parseLevel(value string) int {
	switch value {
	case "*":
		return 1
	case "**":
		return 2
	case "***":
		return 3
	}
	return -1
}

func unmarshalTime(data []byte, t *time.Time) error {
	var err error
	*t, err = time.Parse("2006-01-02 15:04:05", string(data))
	return err
}
