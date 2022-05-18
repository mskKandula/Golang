package main

import (
	"fmt"
)

func main() {

	var (
		intSlice   = []int{1, 2, 4, 5, 6, 7, 8, 9}
		strSlice   = []string{"hello", "world"}
		floatSlice = []float32{1.22, 3.44, 5.66}
	)

	print(intSlice)
	print(strSlice)
	print(floatSlice)

}

func print[T any](slice []T) {
	for _, val := range slice {
		fmt.Printf("%v", val)
	}
	fmt.print("\n")
}
