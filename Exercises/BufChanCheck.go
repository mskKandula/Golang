package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	bufchan = make(chan int, 2)
	wg      sync.WaitGroup
)

func main() {
	fmt.Println("Hello world")
	wg.Add(1)
	go read(bufchan, &wg)
	t1 := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println("time:", time.Since(t1))
		bufchan <- i
	}
	close(bufchan)
	wg.Wait()

}

func read(bufchan <-chan int, wg *sync.WaitGroup) {
	for data := range bufchan {
		fmt.Println(data)
		time.Sleep(4 * time.Second)
	}
	wg.Done()
}
