package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	result := make(chan string, 10)
	count := 0

	go func() {
		words(3, "", result)
		close(result)
	}()

	for range result {
		count += 1
	}

	fmt.Println("Total time taken:", time.Since(t1))
	fmt.Println("Total num of words:", count)
}

func words(length int, prefix string, result chan string) {

	if length < 1 {
		return
	}
	letter := ""
	for i := 97; i < 123; i++ {
		letter = fmt.Sprintf("%s%c", prefix, i)
		//To simulate some processing time
		time.Sleep(1 * time.Millisecond)

		if len(letter) > 2 {
			result <- letter
		}
		//recursive call
		words(length-1, letter, result)
	}
}
