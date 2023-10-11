package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 5, 6, 7, 9, 11, 12}
	k := 2

	fmt.Println(radioTrans(arr, k))
	fmt.Println("8")

}

func radioTrans(arr []int, k int) int {
	count := 1
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > k {
			count += 1
		}
	}
	return count

}
