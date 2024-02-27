package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(superDigit("98963", 100))
	fmt.Println("Time taken---", time.Since(t1))
}

func superDigit(n string, k int) int {
	cache := make(map[int]int)

	if k < 2 {
		num, _ := strconv.Atoi(n)

		return super(num, cache)
	}

	var (
		str string
		sum int
	)

	for i := 0; i < k; i++ {

		if len(str) >= 10 {
			num, _ := strconv.Atoi(str)
			str = ""
			sum += super(num, cache)
		}

		str += n
	}

	if len(str) > 0 {

		num, _ := strconv.Atoi(str)

		sum += super(num, cache)

		if sum > 9 {
			return super(sum, cache)
		}
	}

	return sum
}

func super(n int, cache map[int]int) int {

	if n < 10 {
		return n
	}

	if val, ok := cache[n]; ok {
		return super(val, cache)
	}

	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	cache[n] = sum

	return super(sum, cache)

}
