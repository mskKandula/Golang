package main

import "fmt"

func main() {
	arr := []int{8, 6, 3, 20, 15, 10, 25}
	target := 18
	num1, num2 := twosum(arr, target)
	fmt.Println(num1, num2)
}

func twosum(arr []int, target int) (int, int) {
	if len(arr) < 2 && target < 2 {
		return 0, 0
	}

	dict := make(map[int]int)

	end := len(arr)

	for i := 0; i < end; i++ {
		temp := target - arr[i]

		if _, ok := dict[temp]; ok {
			return dict[temp], i
		}

		dict[arr[i]] = i

	}
	return 0, 0
}
