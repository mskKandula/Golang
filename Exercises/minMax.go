package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)
}

func miniMaxSum(arr []int32) {
	var (
		minVal   int32 = math.MaxInt32
		maxVal   int32 = math.MinInt32
		totalSum int64
	)

	for i := 0; i < len(arr); i++ {
		totalSum += int64(arr[i])
		fmt.Println(totalSum)
		if arr[i] < minVal {
			minVal = arr[i]
		}

		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	fmt.Println(minVal, maxVal, totalSum)

	fmt.Printf("%d %d", totalSum-int64(maxVal), totalSum-int64(minVal))
}
