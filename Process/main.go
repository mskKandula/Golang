package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func odd() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		if i%2 != 0 {
			fmt.Println("Mary")
		}
	}
}

func even() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			fmt.Println("Bob")
		}
	}
}

func main() {
	t1 := time.Now()
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	go odd()
	go even()
	wg.Wait()
	fmt.Println(time.Since(t1))
}
