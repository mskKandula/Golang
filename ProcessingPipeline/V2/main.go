package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t1 := time.Now()

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	ch3 := make(chan int, 5)

	result := make(chan int, 5)

	workers := 3

	go func() {
		wg.Wait()
		close(result)
	}()

	go stage1(ch1, ch2)
	go stage2(ch2, ch3)

	for i := 0; i < workers; i++ {
		go stage3(ch3, result, &wg)
		wg.Add(1)
	}

	for i := 0; i < 7; i++ {
		ch1 <- i
	}

	close(ch1)

	for res := range result {
		fmt.Println("Result is --- ", res)
	}

	fmt.Println(time.Since(t1))
}

func stage1(ch1 <-chan int, ch2 chan<- int) {

	for ch := range ch1 {
		time.Sleep(1 * time.Second)
		ch2 <- ch
	}
	close(ch2)

}

func stage2(ch2 <-chan int, ch3 chan<- int) {

	for ch := range ch2 {
		time.Sleep(1 * time.Second)
		ch3 <- ch
	}
	close(ch3)
}

func stage3(ch3 <-chan int, result chan<- int, wg *sync.WaitGroup) {

	for ch := range ch3 {
		time.Sleep(3 * time.Second)
		result <- ch
	}
	wg.Done()

}
