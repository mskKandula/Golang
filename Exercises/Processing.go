package main

import (
	"fmt"
	"time"
)

func main() {

	batchSize := 3

	toAdd := 5

	resultChan := make(chan int)

	go func() {
		for {
			val := <-resultChan
			fmt.Println(val)
		}
	}()

	for i := 0; i < 21; i++ {

		if i%batchSize == 0 {

			fmt.Println("In Batch", i)

			time.Sleep(3 * time.Second)
		}
		resultChan <- i * toAdd

	}
	close(resultChan)
}
