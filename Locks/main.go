package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int64
	wg    sync.WaitGroup
)

func inc() {
	defer wg.Done()

	count += 1

}

func dec() {
	defer wg.Done()

	count -= 1

}

func main() {
	t1 := time.Now()

	for i := 0; i < 10000000; i++ {
		wg.Add(2)
		go inc()
		go dec()
	}

	wg.Wait()
	fmt.Println(count, time.Since(t1))
}
