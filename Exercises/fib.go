package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(10))
}

func fib(n int) int {
	num1 := 0
	num2 := 1
	res := 0

	for i := 0; i < n; i++ {
		res = num1 + num2
		fmt.Println(res)
		num1 = num2
		num2 = res
	}
	return num1
}
