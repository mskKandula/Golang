package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func Barber(wakeUp <-chan struct{}, buffChan chan int) {

	for {
		select {
		case <-wakeUp:
			process(buffChan)

		default:
			fmt.Println("Barber is sleeping")
		}
	}

}

func process(buffChan chan int) {
	fmt.Println("Barber wokeup & working on the available customers")
	defer wg.Done()

	for val := range buffChan {
		fmt.Println("Processing the customer: ", val)
		time.Sleep(time.Duration(val) * time.Second)
	}
}

func main() {
	fmt.Println("Sleeping Barber")

	buffChan := make(chan int, 5)
	wakeUp := make(chan struct{})

	go Barber(wakeUp, buffChan)
	count := 0

	for {
		if count > 3 {
			break
		}
		count++
		wg.Add(1)

		for i := 0; i < 5; i++ {
			buffChan <- i
		}

		wakeUp <- struct{}{}
	}

	go func() {
		close(buffChan)
		wg.Wait()
	}()

}
