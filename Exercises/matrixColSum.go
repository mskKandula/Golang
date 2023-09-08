package main

import (
	"fmt"
)

func main() {
	var (
		arr       [5][5]int
		resultarr [5]int
	)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			arr[i][j] = i
		}
	}

	for i := 0; i < len(arr); i++ {
		var colArr []int
		for j := 0; j < len(arr); j++ {
			colArr = append(colArr, arr[j][i])
		}
		resultarr[i] = colSum(colArr)
	}

	fmt.Println(resultarr)
}

func colSum(arr []int) int {
	var result int
	for _, val := range arr {
		result += val
	}
	return result
}
