package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func zipFile(file *os.File) error {

	archive, err := os.Create("sample.zip")
	if err != nil {
		return err
	}

	// Create a new zip archive.
	zipWriter := zip.NewWriter(archive)

	writer, err := zipWriter.Create(file.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	err = zipWriter.Close()

	if err != nil {
		return err
	}

	return nil
}

func main() {

	file, err := os.Open("image.png")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// passing file to zip function
	err = zipFile(file)
	if err != nil {
		log.Fatal(err)
	}
}
