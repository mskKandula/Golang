package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {

	zipFile := "../ImageFiles.zip"

	reader, err := zip.OpenReader(zipFile)

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	for _, file := range reader.File {

		fpath := file.Name

		if file.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// Open dest File
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Fatal(err)
		}

		// Open src File
		inFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		// Copy src to dest
		_, err = io.Copy(outFile, inFile)
		if err != nil {
			log.Fatal(err)
		}

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		inFile.Close()
	}

	fmt.Println("Done")
}
