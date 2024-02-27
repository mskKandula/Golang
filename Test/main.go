package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)
	ch1 := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go even(&wg, ch1)
	go odd(&wg, ch1)

	wg.Wait()

}

func odd(wg *sync.WaitGroup, ch1 chan bool) {
	for i := 0; i < 10; i++ {
		<-ch1
		if i%2 != 0 {
			fmt.Println("Odd :", i)
		}
		ch1 <- true
	}
	wg.Done()
}

func even(wg *sync.WaitGroup, ch1 chan bool) {
	for i := 0; i < 10; i++ {
		ch1 <- true
		if i%2 == 0 {
			fmt.Println("Even :", i)
		}
		<-ch1

	}
	wg.Done()
}
