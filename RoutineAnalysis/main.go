package main

import (
	"fmt"
	"runtime"
	"sync"
)

func cpuIntensive(p *int64) {
	defer wg.Done()
	for i := int64(1); i <= 10000000; i++ {
		*p = i
	}
	fmt.Println("Done intensive thing")
}

func printVar(p *int64) {
	defer wg.Done()
	fmt.Printf("print x = %d.\n", *p)
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	x := int64(0)
	wg.Add(2)
	go cpuIntensive(&x) // This should go into background
	go printVar(&x)

	// This won't get scheduled until everything has finished.
	wg.Wait() // Wait for goroutines to finish
}
