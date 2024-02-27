// Number of occurrences of each character in a string

// package main

// import "fmt"

// func main() {

// 	str := "mohana"

// 	dict := make(map[rune]int)

// 	for _, s := range str {

// 		dict[s] += 1
// 	}

// 	for key, val := range dict {
// 		fmt.Println(string(key), "-", val)
// 	}

// }

// sort an integer array

package main

import "fmt"

func main() {

	arr := []int{3, 2, 5, 4, 1}
	size := len(arr)
	fmt.Println(bubbleSort(arr, size))

}

func bubbleSort(arr []int, size int) []int {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(arr)

	}
	return arr

}
