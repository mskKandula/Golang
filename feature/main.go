package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, _ := os.Getwd()

	path := filepath.Join(wd, "temp", "main.txt")
	_, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}
}
