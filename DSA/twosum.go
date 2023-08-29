package main

import "fmt"

func main() {
	arr := []int{3, 6, 8, 10, 15, 20, 25}
	target := 18
	num1, num2 := twosum(arr, target)
	fmt.Println(num1, num2)
}

func twosum(arr []int, target int) (int, int) {
	if len(arr) < 2 && target < 2 {
		return 0, 0
	}

	start, end := 0, len(arr)-1

	for start < end {
		val := arr[start] + arr[end]
		if val > target {
			end -= 1
		} else if val < target {
			start += 1
		} else {
			return arr[start], arr[end]
		}
	}
	return 0, 0
}
