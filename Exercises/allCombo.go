package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{1, 2, 3}

	for i := 0; i < int(math.Pow(2, float64(len(data)))); i++ {
		var arr []int
		for j := 0; j < len(data); j++ {

			if (i & (1 << j)) > 0 {
				arr = append(arr, data[j])
			}

		}
		fmt.Println(arr)
	}
}
