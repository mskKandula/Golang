package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var files []string

// iterate over the files

func visit(files *[]string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// specifying the paths of files which are of pdf format

		if filepath.Ext(path) == ".pdf" {
			*files = append(*files, path)
		}
		return nil

	}

}

func main() {

	// name of the zip file that will be created
	outFile, err := os.Create("ZipExample" + ".zip")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	// Create a new zip archive.
	wr := zip.NewWriter(outFile)

	// Files folder in machine
	rootPath := "pdfFiles"

	err = filepath.Walk(rootPath, visit(&files))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		actualPath := strings.Split(string(file), string("\\"))

		modifiedPath := strings.Join(actualPath[1:], "/")

		// adding filePaths to zip archive
		zipFile, err := wr.Create(modifiedPath)

		if err != nil {
			log.Fatal(err)
		}

		data, _ := ioutil.ReadFile(file)

		// writing the file content at added file path
		_, err = zipFile.Write(data)

		if err != nil {
			log.Fatal(err)
		}

	}
	defer wr.Close()

	if err != nil {
		log.Fatal(err)
	}

}
