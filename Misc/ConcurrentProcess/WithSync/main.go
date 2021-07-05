package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	dateFolders   int = 30
	editorFolders int = 15
	videos        int = 10
	workers       int = 10
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	jobs := make(chan int, 15)

	for i := 0; i < workers; i++ {
		go worker(jobs, &wg)
	}

	for k := 0; k < dateFolders; k++ {
		for l := 0; l < editorFolders; l++ {
			for m := 0; m < videos; m++ {
				jobs <- m
			}
		}
	}

	close(jobs)

	wg.Wait()

	fmt.Println(time.Since(start))

}

func worker(jobs <-chan int, wg *sync.WaitGroup) {

	for job := range jobs {
		wg.Add(1)
		work(job, wg)
	}

}

func work(k int, wg *sync.WaitGroup) {

	fmt.Println("WORKING ON --- ", k)
	time.Sleep(time.Millisecond * (time.Duration(rand.Intn(10) + 1)))
	fmt.Println("FINISHED ---- ", k)
	wg.Done()
}
