package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("No of cpu & goroutines", runtime.NumCPU(), runtime.NumGoroutine())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

		}
	}()
	wg.Add(1)
	go another()
	fmt.Println("No of cpu & goroutines", runtime.NumCPU(), runtime.NumGoroutine())
	wg.Wait()
}

func another() {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("In another")
}
