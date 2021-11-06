package main

import "fmt"

func main() {
	arr := []int{16, 17, 4, 3, 5, 2}
	var j int
	// Leader element is greater/equal to all elements on its right hand side
	for i := 0; i < len(arr); i++ {
		for j = i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				break
			}
		}
		// here 2 is not considered a leader
		if j == len(arr) && i != len(arr)-1 {
			fmt.Printf("---%v", arr[i])
		}
	}
}
