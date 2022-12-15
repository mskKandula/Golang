package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := [][]int{
		{1, 2}, {3, 4}, {5, 6},
	}

	for i := 0; i < len(data)+1; i++ {
		val := fmt.Sprintf("%03b", i)
		j := 0
		for _, b := range val {
			k, _ := strconv.Atoi(string(b))
			fmt.Printf("%d---", data[j][k])
			j++
		}
		fmt.Println("")
	}

	for i := 4; i < 8; i++ {
		val := fmt.Sprintf("%03b", i)
		j := 0
		for _, b := range val {
			k, _ := strconv.Atoi(string(b))
			fmt.Printf("%d---", data[j][k])
			j++
		}
		fmt.Println("")
	}
}
