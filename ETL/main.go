package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jszwec/csvutil"
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

	numRecords, numErrors, err := ETL(file, stmt)

	duration := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}

	frac := float64(numErrors) / float64(numRecords)
	if frac > 0.1 {
		log.Fatalf("too many errors: %d/%d = %f", numErrors, numRecords, frac)
	}
	fmt.Printf("%d records (%.2f errors) in %v\n", numRecords, frac, duration)
}

func ETL(csvFile io.Reader, stmt *sql.Stmt) (int, int, error) {
	r := csv.NewReader(csvFile)
	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		return 0, 0, err
	}
	dec.Register(unmarshalTime)
	numRecords := 0
	numErrors := 0

	for {
		numRecords++
		var row Row
		err = dec.Decode(&row)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error: %d: %s", numRecords, err)
			numErrors++
			continue
		}
		row.Level = parseLevel(row.Viollevel)
		if _, err := stmt.Exec(row); err != nil {
			return 0, 0, err
		}
	}

	return numRecords, numErrors, nil
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
	*t, err = time.Parse("2006-01-02 15:04:05", string(data))
	return err
}
