package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg    sync.WaitGroup
		count int32
	)

	result := make(chan string, 10)
	t1 := time.Now()

	wg.Add(1)
	go words(3, "", result, &wg)

	go func() {
		wg.Wait()
		close(result)
	}()

	for range result {
		count++
	}

	fmt.Println("Total time taken:", time.Since(t1))
	fmt.Println("Total num of words:", count)
}

func words(length int, prefix string, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	if length < 1 {
		return
	}

	letter := ""

	for i := 97; i < 123; i++ {
		letter = fmt.Sprintf("%s%c", prefix, i)
		time.Sleep(1 * time.Millisecond)

		if len(letter) > 2 {
			result <- letter
		}

		wg.Add(1)
		go words(length-1, letter, result, wg)
	}
}
