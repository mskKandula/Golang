package main

import (
	"fmt"
)

func main() {
	prev, curr, t := 0, 1, 1000000000
	for i := 0; i < 55; i++ {
		fmt.Println(prev % t)
		temp := curr
		curr = prev + curr
		prev = temp
	}
}
