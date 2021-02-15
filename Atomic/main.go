package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg  sync.WaitGroup
	sum uint32
)

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&sum, 1) //remove this line & check the sum
		}()

	}
	wg.Wait()
	fmt.Println(sum) //it's 1000 as expected with sync/atomic

}
