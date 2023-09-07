package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		arr       [100][100]int
		resultarr [100]int
	)
	ch := make(chan int)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			arr[i][j] = rand.Intn(10000)
		}
	}
	t1 := time.Now()

	for i := 0; i < len(arr); i++ {
		go func(i int) {
			rowSum(arr[i], ch)
		}(i)
	}

	// Collect the results from Goroutines
	for i := 0; i < len(arr); i++ {
		result := <-ch
		resultarr[i] = result
	}

	close(ch)

	fmt.Println(time.Since(t1))
	fmt.Println(resultarr)
}

func rowSum(arr [100]int, ch chan<- int) {
	var result int
	for _, val := range arr {
		result += val
	}
	ch <- result
}
