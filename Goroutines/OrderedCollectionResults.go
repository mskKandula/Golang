package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	var urls = []string{
		"https://facebook.com",
		"https://apple.com",
		"https://amazon.com",
		"https://netflix.com",
		"https://google.com",
	}
	var wg sync.WaitGroup

	// using slice which is of length same as urls variable
	results := make([]string, len(urls))

	for i, url := range urls {

		wg.Add(1)

		go func(j int, url string) {

			defer wg.Done()

			_, err := http.Get(url)

			if err != nil {
				results[j] = fmt.Sprintf("Error while getting response from %s at index %d", url, j)
			}

			// No race condition occurs, No two goroutines will try to access the same slice index.

			// Storing the result in an ordered fashion
			results[j] = fmt.Sprintf("Success while getting response from %s at index %d", url, j)

		}(i, url)
	}

	wg.Wait()

	for index, val := range results {
		fmt.Printf("%d : %s\n", index, val)
	}
}
