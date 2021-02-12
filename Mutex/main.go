package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	sum int
	mu  sync.Mutex
)

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() //Run Program without Locking & check the result
			sum = sum + 1
			mu.Unlock()

		}()

	}
	wg.Wait()
	fmt.Println(sum) //it's 1000 as expected

}
