package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		arr       [100][100]int
		resultarr [100]int
		wg        sync.WaitGroup
	)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			arr[i][j] = rand.Intn(10000)
		}
	}
	t1 := time.Now()

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resultarr[i] = rowSum(arr[i])
		}(i)
	}

	fmt.Println(time.Since(t1))
	fmt.Println(resultarr)
}

func rowSum(arr [100]int) int {
	var result int
	for _, val := range arr {
		result += val
	}
	return result
}
