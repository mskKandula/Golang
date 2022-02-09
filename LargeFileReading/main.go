package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	mb = 1024 * 1024
	gb = 1024 * mb

	// Change this for actual file path, Take fie of min 1GB & check the performance
	samplefilePath = "Sample.txt"
)

var (
	// Current signifies the counter for bytes of the file.
	current int64

	// Limit signifies the chunk size of file to be processed by every thread.
	limit int64 = 200 * mb
)

func main() {

	wg := sync.WaitGroup{}

	channel := make(chan string)
	dict := make(map[string]int64)
	done := make(chan bool, 1)

	go func() {
		for word := range channel {
			dict[word]++
		}
		done <- true
	}()

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go func() {
			processFile(current, limit, samplefilePath, channel)
			fmt.Printf("%d thread has been completed \n", i)
			wg.Done()
		}()

		current += limit + 1
	}

	// Wait for all go routines to complete.
	wg.Wait()
	close(channel)

	<-done
	close(done)
}

func processFile(offset int64, limit int64, fileName string, channel chan (string)) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Move the pointer of the file to the start of designated chunk.
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)

	// This block of code ensures that the start of chunk is a new word. If
	// a character is encountered at the given position it moves a few bytes till
	// the end of the word.
	if offset != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	var cummulativeSize int64

	for {
		// Break if read size has exceed the chunk size.
		if cummulativeSize > limit {
			break
		}

		b, err := reader.ReadBytes(' ')

		// Break if end of file is encountered.
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		cummulativeSize += int64(len(b))

		s := strings.TrimSpace(string(b))
		if s != "" {
			// Send the read word in the channel to enter into dictionary.
			channel <- s
		}
	}
}
