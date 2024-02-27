package main

import "fmt"

func main() {

	arr := []int{1, 2, 2, 3, 4, 5, 6}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == arr[mid-1] || arr[mid] == arr[mid+1] {
			fmt.Println(arr[mid])
			break
		} else if (arr[start]+arr[end])/2 != arr[mid] {

			start = mid + 1

		} else {

			end = mid - 1

		}

	}

}
