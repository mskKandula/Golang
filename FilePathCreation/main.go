package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	filepath := filepath.Join("../files/text", "golang", "golang.txt")

	file, err := create(filepath)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()
}

// file path creation
func create(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil && !os.IsExist(err) {
		return nil, err
	}
	return os.Create(path)
}
