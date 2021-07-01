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

	totalWork := dateFolders * editorFolders * videos

	//Since Buffered channel is not Blocking
	jobs := make(chan int, totalWork)

	results := make(chan int, totalWork)

	for i := 0; i < workers; i++ {
		go worker(jobs, results)
	}

	for k := 0; k < dateFolders; k++ {
		for l := 0; l < editorFolders; l++ {
			for m := 0; m < videos; m++ {
				jobs <- m
			}
		}
	}

	close(jobs)

	for res := range results {
		fmt.Println("Result is --- ", res)
	}

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
