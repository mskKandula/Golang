package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	_ "runtime"
	"strings"
	"sync"
	"time"
)

var (
	// Files folder in machine
	rootPathPath string = "../Bl's"

	channel = make(chan string, 5)
)

// iterate over the files
func visit() filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		// specifying the paths of files which are of pdf format
		if filepath.Ext(path) == ".pdf" {

			channel <- path

		}
		return nil

	}
}

func zippy(w *zip.Writer, wg *sync.WaitGroup) {

	for file := range channel {

		full := strings.Split(file, string("\\"))

		last := strings.Join(full[1:], "/")

		f, err := w.Create(last)

		if err != nil {
			log.Fatal(err)
		}

		data, _ := ioutil.ReadFile(file)

		_, err = f.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Done()

}

func main() {

	// Anonymous goroutine, closes the channel
	go func() {
		defer close(channel)
		filepath.Walk(rootPathPath, visit())
	}()

	var wg sync.WaitGroup

	t3 := time.Now()

	// name of the zip file that will be created
	outFile, err := os.Create("ZipExample" + "_" + t3.Format("2006-Jan-02") + ".zip")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// adding the goroutine to the waitgroup
	wg.Add(1)

	go zippy(w, &wg)

	wg.Wait()

	defer w.Close()

	fmt.Println(time.Since(t3))

}
