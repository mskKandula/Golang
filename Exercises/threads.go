package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	c  int
	wg sync.WaitGroup
)

func inc() {

	defer wg.Done()
	c++
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Hello")

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go inc()
	}
	wg.Wait()
	fmt.Println(c)
}
