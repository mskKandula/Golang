package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ch1 := make(chan bool)
	ch1 := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go even(&wg, ch1)
	go odd(&wg, ch1)

	ch1 <- true
	wg.Wait()
	close(ch1)

}

func odd(wg *sync.WaitGroup, ch1 chan bool) {
	for i := 1; i < 10; i = i + 2 {
		<-ch1
		time.Sleep(1 * time.Second)
		fmt.Println("Odd :", i)

		ch1 <- true
	}
	// close(ch1)
	wg.Done()
}

func even(wg *sync.WaitGroup, ch1 chan bool) {
	for i := 2; i < 10; i = i + 2 {
		<-ch1
		time.Sleep(1 * time.Second)
		fmt.Println("Even :", i)

		ch1 <- true

	}
	// close(ch1)
	wg.Done()
}
