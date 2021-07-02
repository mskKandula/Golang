package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	dateFolders   int = 30
	editorFolders int = 15
	videos        int = 10
	workers       int = 10
)

func main() {
	start := time.Now()

	jobs := make(chan int, 15)

	results := make(chan int, 15)

	done := make(chan struct{})

	for i := 0; i < workers; i++ {
		go worker(jobs, results)
	}

	//Channel Reader must be Initialised before channel writer
	go func() {
		for res := range results {
			fmt.Println("Result is --- ", res)
		}
		done <- struct{}{}
	}()

	for k := 0; k < dateFolders; k++ {
		for l := 0; l < editorFolders; l++ {
			for m := 0; m < videos; m++ {
				jobs <- m
			}
		}
	}

	close(jobs)

	<-done

	fmt.Println(time.Since(start))

}

func worker(jobs <-chan int, results chan<- int) {

	for job := range jobs {

		results <- work(job)
	}
	close(results)

}

func work(k int) int {

	fmt.Println("WORKING ON --- ", k)
	time.Sleep(time.Millisecond * (time.Duration(rand.Intn(10) + 1)))
	fmt.Println("FINISHED ---- ", k)
	return k
}
