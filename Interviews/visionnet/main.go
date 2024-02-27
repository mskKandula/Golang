package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i = i + 2 {
			fmt.Println("first Goroutine:", i)
		}

	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("second Goroutine:", i)
		}

	}()

	wg.Wait()

}

//items

//post api/v1/items

//get api/


login ---> JWt ----> response 

