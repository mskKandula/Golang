package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("boston.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	_, _, err = ETL(file)

	duration := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Time Taken to Process: ", duration)
}

func ETL(csvFile io.Reader) (int, int, error) {

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
