package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	cache := make(map[int]int)
	fmt.Println(fact(10, cache), time.Since((t1)))

}

func fact(n int, cache map[int]int) int {

	if n < 2 {
		return n
	}

	if val, ok := cache[n]; ok {
		return val
	}

	res := fact(n-1, cache) + fact(n-2, cache)
	cache[n] = res
	return res
}
